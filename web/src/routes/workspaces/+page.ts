import type { Workspace } from "$lib/api/types";
import { workspaces } from "$lib/api/client";

export interface PageData {
	workspaces: Workspace[];
	error?: string;
}

export async function load(): Promise<PageData> {
	const response = await workspaces.list();

	if (response.success && response.data) {
		return { workspaces: response.data };
	}

	return {
		workspaces: [],
		error: response.error || response.message || "Failed to load workspaces",
	};
}
