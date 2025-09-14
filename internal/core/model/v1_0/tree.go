package v1_0

import "github.com/samber/lo"

type Node struct {
	Content  *ContentObject
	Parent   *Node
	Children []*Node
	Level    int
}

type Tree struct {
	Root *Node
}

type TreeBuilder struct {
	MaxLevel  int
	OpenLinks bool
}

func (b *TreeBuilder) Build(source *ERPN, dig bool) *Tree {
	result := &Tree{}

	source.Tree = result

	result.Root = b.buildNode(source.Content, nil, 0, dig)

	return result
}

func (b *TreeBuilder) buildNode(content *ContentObject, from *Node, level int, dig bool) *Node {
	result := &Node{
		Content:  content,
		Parent:   from,
		Children: nil,
		Level:    level,
	}

	content.Node = result

	if !dig {
		return result
	}

	if b.MaxLevel != 0 && level > b.MaxLevel {
		return result
	}

	result.Children = make([]*Node, len(content.Content))
	for i, object := range content.Content {
		if object == nil {
			continue
		}

		if object.Link != nil {
			if !b.OpenLinks {
				continue
			}

			// TODO add opening link
			continue
		}

		result.Children[i] = b.buildNode(object, result, level+1, dig)
	}

	interlevel := b.buildResourses(result.Children)
	if interlevel == nil {
		return result
	}

	content.Essence.Realization.Value = b.mergeResources(content.Essence.Realization.Value, interlevel.Value)
	content.Essence.Realization.Input = b.mergeResources(content.Essence.Realization.Input, interlevel.Input)
	content.Essence.Realization.Output = b.mergeResources(content.Essence.Realization.Output, interlevel.Output)
	content.Essence.Realization.Resources = b.mergeResources(content.Essence.Realization.Resources, interlevel.Resources)
	content.Essence.Realization.Value = b.mergeResources(content.Essence.Realization.Value, interlevel.Value)

	return result
}

func (b *TreeBuilder) buildResourses(pool []*Node) *Interlevel {
	// TODO add counting (count to extra, left also go up)
	result := &Interlevel{}
	inPool := make(map[string]*CommonObject, len(pool))
	outPool := make(map[string]*CommonObject, len(pool))

	for _, node := range pool {
		if node == nil {
			continue
		}

		if node.Content == nil {
			continue
		}

		result.Resources = b.mergeResources(result.Resources, node.Content.Essence.Realization.Resources)
		b.hydrateEnd(result.Resources, node)

		result.Value = b.mergeResources(result.Value, node.Content.Essence.Realization.Value)
		b.hydrateStart(result.Value, node)

		inPool = b.mergeMap(inPool, node.Content.Essence.Realization.Input)
		b.hydrateEnd(node.Content.Essence.Realization.Input, node)

		outPool = b.mergeMap(outPool, node.Content.Essence.Realization.Output)
		b.hydrateStart(node.Content.Essence.Realization.Output, node)
	}

	result.Input = lo.Values(lo.OmitByKeys(inPool, lo.Keys(outPool)))
	result.Output = lo.Values(lo.OmitByKeys(outPool, lo.Keys(inPool)))

	return nil
}

func (b *TreeBuilder) mergeResources(appendTo, extra []*CommonObject) []*CommonObject {
	return lo.UniqBy(append(appendTo, extra...), func(item *CommonObject) string { return *item.Name })
}

func (b *TreeBuilder) mergeMap(appendTo map[string]*CommonObject, extra []*CommonObject) map[string]*CommonObject {
	uniq := lo.UniqBy(extra, func(item *CommonObject) string { return *item.Name })
	for i, object := range uniq {
		appendTo[*object.Name] = extra[i]
	}

	for i, object := range extra {
		extra[i] = appendTo[*object.Name]
	}

	return appendTo
}

func (b *TreeBuilder) hydrateStart(targets []*CommonObject, node *Node) {
	for _, target := range targets {
		target.start = append(target.start, node)
	}
}

func (b *TreeBuilder) hydrateEnd(targets []*CommonObject, node *Node) {
	for _, target := range targets {
		target.start = append(target.end, node)
	}
}
