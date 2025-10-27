package example

#Config: {
	// Service name (required, lowercase alphanumeric with hyphens)
	name: string & =~"^[a-z][a-z0-9-]*$"
	
	// Port number (required, 1-65535)
	port: int & >=1 & <=65535
	
	// Enable or disable the service
	enabled: bool
	
	// Deployment environment (required)
	environment: "development" | "staging" | "production"
	
	// Number of replicas (optional, 1-10)
	replicas?: int & >=1 & <=10
	
	// Resource limits (optional)
	resources?: {
		cpu: string & =~"^[0-9]+m$"
		memory: string & =~"^[0-9]+(Mi|Gi)$"
	}
}
