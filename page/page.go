package page

import "github.com/sdvdxl/go-tools/number"

// Page 分页信息
type Page struct {
	Page         int         `json:"page"`
	Size         int         `json:"size"`
	TotalRecords int         `json:"totalRecords"`
	TotalPages   int         `json:"totalPages"`
	Items        interface{} `json:"items"`
}

const (
	// DefaultPageSize 默认分页大小
	DefaultPageSize int = 10

	// MaxPageSize 最大分页大小
	MaxPageSize int = 50
)

// Offset 用于pg 分页参数偏移
func (p Page) Offset() int {
	return p.Size * (p.Page - 1)
}

// PagingPage 从分页对象中获取当前页和页码
func PagingPage(p Page, totalRecords int, items interface{}) Page {
	return Paging(p.Page, p.Size, totalRecords, items)
}

// Paging 分页
func (p Page) Paging() Page {
	return PagingPage(p, p.TotalRecords, p.Items)
}

// Paging 产生分页信息
func Paging(curPage int, pageSize int, totalRecords int, items interface{}) Page {

	totalPages := int(totalRecords) / pageSize
	if totalPages == 0 {
		totalPages = 1
	}

	if int(totalRecords)%pageSize != 0 {
		totalPages++
	}

	return Page{
		Page:         curPage,
		Size:         pageSize,
		TotalRecords: int(totalRecords),
		TotalPages:   totalPages,
		Items:        items,
	}
}

// New 新建一个Fixed Page对象
func New(page, size int) Page {
	if page < 0 {
		page = 1
	}

	if size <= 0 || size > MaxPageSize {
		size = DefaultPageSize
	}

	return Page{Page: page, Size: size}
}

// NewFromString 从字符串中生成一个新的page
// cur 当前页，从1开始
// size 分页大小
// 如果cur和size int 转换失败，分别转换成 defalult
func NewFromString(cur, size string) Page {
	return New(number.DefaultInt(cur, 1), number.DefaultInt(size, DefaultPageSize))
}

// NewDefaultFromString 从字符串中生成一个新的page
// cur 当前页，从1开始
// size 分页大小
// 如果cur和size int 转换失败，分别转换成 defalult
func NewDefaultFromString(cur, size string, curDefault, sizeDefault int) Page {
	return New(number.DefaultInt(cur, curDefault), number.DefaultInt(size, sizeDefault))
}
