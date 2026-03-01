export interface ApiResponse<T> {
	success: boolean;
	data?: T;
	error?: string;
	message?: string;
}

export interface ApiError {
	error: string;
	message: string;
}

export interface Workspace {
	id: string;
	name: string;
	created_at: string;
	updated_at: string;
}

export interface Connection {
	id: string;
	workspace_id: string;
	name: string;
	type: string;
	config: Record<string, unknown>;
	created_at: string;
	updated_at: string;
}
