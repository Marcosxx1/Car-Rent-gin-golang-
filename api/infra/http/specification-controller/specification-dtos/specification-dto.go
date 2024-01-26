package specificationdtos

/* also validate */
type SpecificationInputDto struct {
	Name        string `json:"name" binding:"required"`
	Description string `json:"description"`
	CarID       string `json:"car_id"`
}

type SpecificationOutputDto struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	CarID       string `json:"car_id"`
}
