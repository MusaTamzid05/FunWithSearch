package lib


type Selector interface {
    IsSelected(node Node) bool
}

type SelectorByTextSize struct {
    MinSize int
    MaxSize int
    Key string
}

func (s SelectorByTextSize) IsSelected(node Node) bool {
    textLength := len(node.Data[s.Key])

    if textLength > s.MinSize && textLength <= s.MaxSize {
        return true
    }

    return false

}

func MakeSelectorByTextSize(minSize, maxSize int, key string) SelectorByTextSize {
    return SelectorByTextSize{MinSize: minSize, MaxSize: maxSize, Key: key}
}
