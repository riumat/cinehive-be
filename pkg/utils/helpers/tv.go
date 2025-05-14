package helpers

import (
	"strings"

	"github.com/riumat/cinehive-be/pkg/utils/types"
)

func ExtractCastItems(cast []any) []types.CastItem {
	var castItems []types.CastItem

	for _, item := range cast {
		castMap, ok := item.(map[string]any)
		if !ok {
			continue
		}

		castItem := types.CastItem{
			ID:          getFloat64Value(castMap, "id"),
			Name:        getStringValue(castMap, "name"),
			ProfilePath: getStringValue(castMap, "profile_path"),
			Character:   extractCharacterFromRoles(castMap),
		}

		castItems = append(castItems, castItem)
	}

	return castItems
}

func getFloat64Value(m map[string]any, key string) float64 {
	if value, ok := m[key].(float64); ok {
		return value
	}
	return 0
}

func getStringValue(m map[string]any, key string) string {
	if value, ok := m[key].(string); ok {
		return value
	}
	return ""
}

func extractCharacterFromRoles(castMap map[string]any) string {
	roles, ok := castMap["roles"].([]any)
	if !ok {
		return ""
	}

	var characters []string
	for _, role := range roles {
		roleMap, ok := role.(map[string]any)
		if !ok {
			continue
		}

		if character, ok := roleMap["character"].(string); ok && character != "" {
			characters = append(characters, character)
		}
	}

	return strings.Join(characters, ", ")
}
