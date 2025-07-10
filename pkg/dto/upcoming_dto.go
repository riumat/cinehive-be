package dto

import (
	"github.com/riumat/cinehive-be/pkg/utils"
	"github.com/riumat/cinehive-be/pkg/utils/types"
)

type UpcomingDto struct {
	Movies []types.ContentCard `json:"movies"`
	Tvs    []types.ContentCard `json:"tvs"`
}

type UpcomingResponse = utils.Response[[]types.ContentCard]
