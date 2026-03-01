import type { ApiResponse, ApiError, Workspace, Connection } from './types';

const BASE_URL = '/api';

async function request<T>(
	url: string,
	method: string,
	body?: unknown
): Promise<ApiResponse<T>> {
	const headers: Record<string, string> = {
		'Content-Type': 'application/json',
	};

	try {
		const response = await fetch(`${BASE_URL}${url}`, {
			method,
			headers,
			body: body ? JSON.stringify(body) : undefined,
		});

		const data = await response.json();

		if (response.ok) {
			return data;
		} else {
			return {
				success: false,
				error: data.error || 'Unknown error',
				message: data.message || 'Request failed',
			};
		}
	} catch (error) {
		return {
			success: false,
			error: 'Network error',
			message: error instanceof Error ? error.message : 'Unknown error',
		};
	}
}

export const api = {
	get: <T>(url: string) => request<T>(url, 'GET'),

	post: <T>(url: string, data: unknown) => request<T>(url, 'POST', data),

	put: <T>(url: string, data: unknown) => request<T>(url, 'PUT', data),

	del: <T>(url: string) => request<T>(url, 'DELETE'),
};

export const workspaces = {
	list: (): Promise<ApiResponse<Workspace[]>> =>
		api.get<Workspace[]>('/workspaces'),

	create: (data: { name: string }): Promise<ApiResponse<Workspace>> =>
		api.post<Workspace>('/workspaces', data),
};

export const connections = {
	list: (workspaceId: string): Promise<ApiResponse<Connection[]>> =>
		api.get<Connection[]>(`/workspaces/${workspaceId}/connections`),

	create: (
		workspaceId: string,
		data: { name: string; type: string; config: Record<string, unknown> }
	): Promise<ApiResponse<Connection>> =>
		api.post<Connection>(`/workspaces/${workspaceId}/connections`, data),

	test: (id: string): Promise<ApiResponse<{ connected: boolean }>> =>
		api.post<{ connected: boolean }>(`/connections/${id}/test`, {}),
};
