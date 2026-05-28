package analyzer

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/openai/openai-go/v3"
	"github.com/openai/openai-go/v3/option"
	"github.com/openai/openai-go/v3/responses"
	"github.com/volodymyr-miretskyi/go-nutrition/internal/modules/food/domain"
	"github.com/volodymyr-miretskyi/go-nutrition/internal/modules/food/usecase"
)

type Analyzer struct {
	client openai.Client
}

func New(apiKey string) *Analyzer {
	return &Analyzer{
		client: openai.NewClient(option.WithAPIKey(apiKey)),
	}
}

func (a *Analyzer) AnalyzeFood(ctx context.Context, params *usecase.FoodAnalyzerAnalyzeFoodParams) (*domain.Nutrients, error) {
	prompt := fmt.Sprintf(`
		Analyze this food image URL: %s

		User comment: %s

		Return ONLY JSON:

		{
			"calories": int(value),
			"proteins": int(value),
			"fats": int(value),
			"carbs": int(value)
		}
	`, params.ImageUrl, params.Comment)

	resp, err := a.client.Responses.New(ctx, responses.ResponseNewParams{
		Model: openai.ChatModelGPT4_1Mini,
		Input: responses.ResponseNewParamsInputUnion{
			OfString: openai.String(prompt),
		},
	})

	if err != nil {
		return nil, err
	}

	var parsed *domain.Nutrients

	output := resp.OutputText()

	output = strings.TrimSpace(output)
	output = strings.TrimPrefix(output, "```json")
	output = strings.TrimPrefix(output, "```")
	output = strings.TrimSuffix(output, "```")
	output = strings.TrimSpace(output)

	if err := json.Unmarshal([]byte(output), &parsed); err != nil {
		return nil, err
	}

	return parsed, nil
}
