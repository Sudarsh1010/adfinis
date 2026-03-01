import type { Workspace } from '../api/types';

/**
 * WorkspaceStore - Svelte 5 reactive state for workspace management
 * Uses $state runes (Svelte 5) instead of writable/readable stores (Svelte 4)
 */
export class WorkspaceStore {
  /** Current workspace (null initially) */
  currentWorkspace = $state<Workspace | null>(null);

  /** All workspaces list */
  workspaces = $state<Workspace[]>([]);

  /** Loading state */
  loading = $state(false);

  /** Error message or null if no error */
  error = $state<string | null>(null);

  /**
   * Fetch all workspaces
   * TODO: Implement API call in later iteration
   */
  async fetchWorkspaces(): Promise<void> {
    this.loading = true;
    this.error = null;
    try {
      // TODO: Replace with actual API call
      // const response = await fetch('/api/workspaces');
      // this.workspaces = await response.json();
    } catch (err) {
      this.error = err instanceof Error ? err.message : 'Failed to fetch workspaces';
      throw err;
    } finally {
      this.loading = false;
    }
  }

  /**
   * Create a new workspace
   * @param data - Workspace data (name, description, etc.)
   * @returns Created workspace
   * TODO: Implement API call in later iteration
   */
  async createWorkspace(data: Partial<Workspace>): Promise<Workspace> {
    this.loading = true;
    this.error = null;
    try {
      // TODO: Replace with actual API call
      // const response = await fetch('/api/workspaces', {
      //   method: 'POST',
      //   headers: { 'Content-Type': 'application/json' },
      //   body: JSON.stringify(data),
      // });
      // const newWorkspace = await response.json();
      // this.workspaces = [...this.workspaces, newWorkspace];
      // this.currentWorkspace = newWorkspace;
      // return newWorkspace;

      // Mock implementation for Wave 1
const newWorkspace: Workspace = {
        id: crypto.randomUUID(),
        name: data.name || 'New Workspace',
        created_at: new Date().toISOString(),
        updated_at: new Date().toISOString(),
      };
      return newWorkspace;
    } catch (err) {
      this.error = err instanceof Error ? err.message : 'Failed to create workspace';
      throw err;
    } finally {
      this.loading = false;
    }
  }

  /**
   * Select a workspace as current
   * @param workspaceId - ID of workspace to select
   */
  selectWorkspace(workspaceId: string): void {
    this.currentWorkspace = this.workspaces.find((w) => w.id === workspaceId) || null;
  }
}

/** Global workspace store instance */
export const workspaceStore = new WorkspaceStore();
