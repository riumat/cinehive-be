package dto

import (
	"github.com/riumat/cinehive-be/pkg/utils"
	"github.com/riumat/cinehive-be/pkg/utils/types"
)

type SearchDto = utils.PaginatedResponse[[]types.SearchResult]
