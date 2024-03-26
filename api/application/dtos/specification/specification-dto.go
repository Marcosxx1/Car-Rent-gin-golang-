package specificationdtos

/* also validate */
type SpecificationInputDto struct {
	Name        string `json:"name" validate:"required"`
	Description string `json:"description" validate:"required"`
	CarID       string `json:"car_id" validate:"required"`
}

type SpecificationOutputDto struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	CarID       string `json:"car_id"`
}
