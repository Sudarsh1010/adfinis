import type { Connection } from '../api/types';

/**
 * Uses $state runes (Svelte 5) instead of writable/readable stores (Svelte 4)
 */
export class ConnectionStore {
  /** All connections for the current workspace */
  connections = $state<Connection[]>([]);

  /** Loading state */
  loading = $state(false);

  /** Error message or null if no error */
  error = $state<string | null>(null);

  /**
   * Fetch connections for a specific workspace
   * @param workspaceId - ID of workspace to fetch connections for
   * TODO: Implement API call in later iteration
   */
  async fetchConnections(workspaceId: string): Promise<void> {
    this.loading = true;
    this.error = null;
    try {
      // TODO: Replace with actual API call
      // const response = await fetch(`/api/workspaces/${workspaceId}/connections`);
      // this.connections = await response.json();
    } catch (err) {
      this.error = err instanceof Error ? err.message : 'Failed to fetch connections';
      throw err;
    } finally {
      this.loading = false;
    }
  }

  /**
   * Create a new connection
   * @param workspaceId - ID of workspace to create connection in
   * @param data - Connection data (name, type, url, credentials)
   * @returns Created connection
   * TODO: Implement API call in later iteration
   */
  async createConnection(workspaceId: string, data: Partial<{ name?: string; type?: string; config?: { url?: string; username?: string; apiKey?: string; databaseName?: string; testResult?: 'success' | 'failure' | 'pending' } }>): Promise<Connection> {
    this.loading = true;
    this.error = null;
    try {
      // TODO: Replace with actual API call
      // const response = await fetch(`/api/workspaces/${workspaceId}/connections`, {
      //   method: 'POST',
      //   headers: { 'Content-Type': 'application/json' },
      //   body: JSON.stringify(data),
      // });
      // const newConnection = await response.json();
      // this.connections = [...this.connections, newConnection];
      // return newConnection;

      // Mock implementation for Wave 1
      const now = new Date().toISOString();
      const newConnection: Connection = {
        id: crypto.randomUUID(),
        workspace_id: '',
        name: data.name || 'New Connection',
        type: data.type || 'api',
        config: {
          url: data.config?.url || '',
          username: data.config?.username,
          apiKey: data.config?.apiKey,
          databaseName: data.config?.databaseName,
          testResult: 'pending'
        },
        created_at: now,
        updated_at: now,
      };
      return newConnection;
    } catch (err) {
      this.error = err instanceof Error ? err.message : 'Failed to create connection';
      throw err;
    } finally {
      this.loading = false;
    }
  }

  /**
   * Test a connection
   * @param id - ID of connection to test
   * @returns Test result status
   * TODO: Implement actual connection testing in later iteration
   */
  async testConnection(id: string): Promise<'success' | 'failure'> {
    this.loading = true;
    this.error = null;
    try {
      // TODO: Replace with actual API call
      // const response = await fetch(`/api/connections/${id}/test`, {
      //   method: 'POST',
      // });
      // const result = await response.json();
      // this.connections = this.connections.map(c =>
      //   c.id === id ? { ...c, testResult: result.success ? 'success' : 'failure' } : c
      // );
      // return result.success ? 'success' : 'failure';

      // Mock implementation for Wave 1
      await new Promise((resolve) => setTimeout(resolve, 500)); // Simulate network delay
      const connection = this.connections.find((c) => c.id === id);
      if (!connection) {
        throw new Error('Connection not found');
      }

      // Mock test result - always success for Wave 1
      this.connections = this.connections.map((c) =>
        c.id === id ? { ...c, config: { ...c.config, testResult: 'success' } } : c
      );
      return 'success';
    } catch (err) {
      this.error = err instanceof Error ? err.message : 'Failed to test connection';
      this.connections = this.connections.map((c) =>
        c.id === id ? { ...c, config: { ...c.config, testResult: 'failure' } } : c
      );
      return 'failure';
    } finally {
      this.loading = false;
    }
  }

  /**
   * Delete a connection
   * @param id - ID of connection to delete
   */
  deleteConnection(id: string): void {
    this.connections = this.connections.filter((c) => c.id !== id);
  }

  /**
   * Update a connection
   * @param id - ID of connection to update
   * @param data - Partial connection data
   */
  updateConnection(id: string, data: Partial<Connection>): void {
    this.connections = this.connections.map((c) =>
      c.id === id ? { ...c, ...data, updated_at: new Date().toISOString() } : c
    );
  }
}

/** Global connection store instance */
export const connectionStore = new ConnectionStore();
