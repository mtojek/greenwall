package healthcheck

// Healthcheck is responsible for storing latest statuses of monitored nodes.
type Healthcheck struct{}

// NewHealthcheck method creates a new instance of healthcheck.
func NewHealthcheck() *Healthcheck {
	return &Healthcheck{}
}

// Status method returns a report containing statuses of monitored nodes.
func (h *Healthcheck) Status() HealthStatus {
	//TODO

	frontendGroup := Group{
		Name:   "Frontend Nodes (us-east-1)",
		Anchor: "frontend_nodes__us-east-1_",
		Nodes: []Node{
			{
				Name:     "front-1",
				Endpoint: "https://www.example.com/",
				Status:   "success",
				Message:  "OK",
			},
			{
				Name:     "front-2",
				Endpoint: "https://www.example.com/",
				Status:   "success",
				Message:  "OK",
			},
			{
				Name:     "front-3",
				Endpoint: "https://www.example.com/",
				Status:   "danger",
				Message:  "Something went wrong, really really bad",
			},
		},
	}

	middlewareGroup := Group{
		Name:   "Middleware Nodes (us-west-2)",
		Anchor: "middleware_nodes__us-west-2_",
		Nodes: []Node{
			{
				Name:     "middleware-1",
				Endpoint: "https://www.example.com/",
				Status:   "success",
				Message:  "OK",
			},
			{
				Name:     "middleware-2",
				Endpoint: "https://www.example.com/",
				Status:   "success",
				Message:  "OK",
			},
			{
				Name:     "middleware-3",
				Endpoint: "https://www.example.com/",
				Status:   "success",
				Message:  "OK",
			},
			{
				Name:     "middleware-4",
				Endpoint: "https://www.example.com/",
				Status:   "success",
				Message:  "OK",
			},
			{
				Name:     "middleware-5",
				Endpoint: "https://www.example.com/",
				Status:   "success",
				Message:  "OK",
			},
		},
	}

	backendGroup := Group{
		Name:   "Backend Nodes (us-east-1)",
		Anchor: "backend_nodes__us-east-1_",
		Nodes: []Node{
			{
				Name:     "backend-1",
				Endpoint: "https://www.example.com/",
				Status:   "success",
				Message:  "OK",
			},
			{
				Name:     "backend-2",
				Endpoint: "https://www.example.com/",
				Status:   "success",
				Message:  "OK",
			},
			{
				Name:     "backend-3",
				Endpoint: "https://www.example.com/",
				Status:   "danger",
				Message:  "Somethin went wrong",
			},
		},
	}

	return HealthStatus{
		Groups: []Group{
			frontendGroup,
			middlewareGroup,
			backendGroup,
		},
	}
}
