package enums

type MaintenanceStatus string

const (
	Scheduled          MaintenanceStatus = "Scheduled"
	InProgress         MaintenanceStatus = "In Progress"
	Completed          MaintenanceStatus = "Completed"
	PendingApproval    MaintenanceStatus = "Pending Approval"
	Canceled           MaintenanceStatus = "Canceled"
	AwaitingParts      MaintenanceStatus = "Awaiting Parts"
	AwaitingPayment    MaintenanceStatus = "Awaiting Payment"
	Rescheduled        MaintenanceStatus = "Rescheduled"
	MaintenanceFailed  MaintenanceStatus = "Maintenance Failed"
	AwaitingInspection MaintenanceStatus = "Awaiting Inspection"
)
   