<script lang="ts">
	import { page } from '$app/stores';
	import { toast } from 'svelte-sonner';
	import type { Connection, Workspace } from '$lib/api/types';
	import { connections, workspaces } from '$lib/api/client';
	import * as Card from '$components/ui/card/index.js';
	import * as Dialog from '$components/ui/dialog/index.js';
	import * as Sidebar from '$components/ui/sidebar/index.js';
	import * as Breadcrumb from '$components/ui/breadcrumb/index.js';
	import * as Select from '$components/ui/select/index.js';
	import { Separator } from '$components/ui/separator/index.js';
	import Button from '$components/ui/button/button.svelte';
	import Input from '$components/ui/input/input.svelte';
	import Label from '$components/ui/label/label.svelte';
	import { Badge } from '$components/ui/badge/index.js';
	import Skeleton from '$components/ui/skeleton/skeleton.svelte';
	import AppSidebar from '$components/app-sidebar.svelte';
	import PlusIcon from '@lucide/svelte/icons/plus';
	import Loader2Icon from '@lucide/svelte/icons/loader-2';
	import CheckCircleIcon from '@lucide/svelte/icons/check-circle';
	import XCircleIcon from '@lucide/svelte/icons/x-circle';
	import AlertCircleIcon from '@lucide/svelte/icons/alert-circle';
	import DatabaseIcon from '@lucide/svelte/icons/database';

	// State using Svelte 5 runes
	let connectionList = $state<Connection[]>([]);
	let workspace = $state<Workspace | null>(null);
	let isLoading = $state(true);
	let error = $state<string | null>(null);

	// Dialog states
	let addDialogOpen = $state(false);
	let isSubmitting = $state(false);
	let isTesting = $state(false);

	// Form states
	let formData = $state({
		name: '',
		type: 'api' as 'api' | 'database',
		url: '',
		apiKey: '',
		username: '',
		password: '',
		databaseName: ''
	});

	// Connection types for select
	const connectionTypes = [
		{ value: 'api', label: 'API / REST' },
		{ value: 'database', label: 'Database' }
	];

	// Get workspace ID from URL
	let workspaceId = $derived($page.params.id ?? '');

	// Load workspace and connections on mount
	$effect(() => {
		if (workspaceId) {
			loadData();
		}
	});

	async function loadData() {
		isLoading = true;
		error = null;

		// Load workspace info
		const workspaceResponse = await workspaces.list();
		if (workspaceResponse.success && workspaceResponse.data) {
			workspace = workspaceResponse.data.find(w => w.id === workspaceId) || null;
		}

		// Load connections
		const connectionsResponse = await connections.list(workspaceId);
		if (connectionsResponse.success && connectionsResponse.data) {
			connectionList = connectionsResponse.data;
		} else {
			error = connectionsResponse.error || 'Failed to load connections';
		}

		isLoading = false;
	}

	function resetForm() {
		formData = {
			name: '',
			type: 'api',
			url: '',
			apiKey: '',
			username: '',
			password: '',
			databaseName: ''
		};
	}

	async function handleTestConnection() {
		if (!formData.url.trim()) {
			toast.error('Please enter a URL to test');
			return;
		}

		isTesting = true;
		try {
			// Simulate test - in real app, this would call an API endpoint
			await new Promise(resolve => setTimeout(resolve, 1500));
			toast.success('Connection test successful');
		} catch (err) {
			toast.error('Connection test failed');
		} finally {
			isTesting = false;
		}
	}

	async function handleSubmit() {
		if (!formData.name.trim() || !formData.url.trim()) {
			toast.error('Please fill in required fields');
			return;
		}

		isSubmitting = true;

		const config: Record<string, unknown> = {
			url: formData.url,
			testResult: 'pending'
		};

		if (formData.type === 'api') {
			config.apiKey = formData.apiKey;
		} else {
			config.username = formData.username;
			config.password = formData.password;
			config.databaseName = formData.databaseName;
		}

		const response = await connections.create(workspaceId, {
			name: formData.name,
			type: formData.type,
			config
		});

		if (response.success && response.data) {
			connectionList = [...connectionList, response.data];
			toast.success('Connection created successfully');
			resetForm();
			addDialogOpen = false;
		} else {
			toast.error(response.error || 'Failed to create connection');
		}

		isSubmitting = false;
	}

	async function handleTest(id: string) {
		toast.info('Testing connection...');
		const response = await connections.test(id);
		if (response.success && response.data?.connected) {
			toast.success('Connection test successful');
		} else {
			toast.error('Connection test failed');
		}
	}

	function getStatusBadge(testResult?: string) {
		switch (testResult) {
			case 'success':
				return { variant: 'default' as const, icon: CheckCircleIcon, label: 'Connected' };
			case 'failure':
				return { variant: 'destructive' as const, icon: XCircleIcon, label: 'Failed' };
			default:
				return { variant: 'outline' as const, icon: AlertCircleIcon, label: 'Pending' };
		}
	}

	function getConnectionUrl(connection: Connection): string {
		if (connection.config && typeof connection.config === 'object' && 'url' in connection.config) {
			return connection.config.url as string;
		}
		return 'N/A';
	}

	function getTestResult(connection: Connection): string | undefined {
		if (connection.config && typeof connection.config === 'object' && 'testResult' in connection.config) {
			return connection.config.testResult as string;
		}
		return undefined;
	}
</script>

<Sidebar.Provider>
	<AppSidebar />
	<Sidebar.Inset>
		<header
			class="flex h-16 shrink-0 items-center gap-2 transition-[width,height] ease-linear group-has-data-[collapsible=icon]/sidebar-wrapper:h-12"
		>
			<div class="flex items-center gap-2 px-4">
				<Sidebar.Trigger class="-ms-1" />
				<Separator orientation="vertical" class="me-2 data-[orientation=vertical]:h-4" />
				<Breadcrumb.Root>
					<Breadcrumb.List>
						<Breadcrumb.Item>
							<Breadcrumb.Link href="/settings">Settings</Breadcrumb.Link>
						</Breadcrumb.Item>
						<Breadcrumb.Separator />
						<Breadcrumb.Item>
							<Breadcrumb.Link href="/settings/workspaces">Workspaces</Breadcrumb.Link>
						</Breadcrumb.Item>
						<Breadcrumb.Separator />
						<Breadcrumb.Item>
							<Breadcrumb.Link href="/settings/workspaces/{workspaceId}">
								{workspace?.name ?? 'Workspace'}
							</Breadcrumb.Link>
						</Breadcrumb.Item>
						<Breadcrumb.Separator />
						<Breadcrumb.Item>
							<Breadcrumb.Page>Connections</Breadcrumb.Page>
						</Breadcrumb.Item>
					</Breadcrumb.List>
				</Breadcrumb.Root>
			</div>
		</header>

		<div class="flex flex-1 flex-col gap-4 p-4 pt-0">
			<!-- Error Alert -->
			{#if error}
				<div
					class="bg-destructive/10 border-destructive/20 text-destructive flex items-center gap-2 rounded-lg border p-4"
				>
					<AlertCircleIcon class="size-5 shrink-0" />
					<span>{error}</span>
					<Button variant="ghost" size="sm" class="ml-auto" onclick={() => (error = null)}>
						Dismiss
					</Button>
				</div>
			{/if}

			<!-- Main Content Card -->
			<Card.Root>
				<Card.Header>
					<div class="flex items-center justify-between">
						<div>
							<Card.Title>Connections</Card.Title>
							<p class="text-muted-foreground mt-1 text-sm">
								Manage data source connections for {workspace?.name ?? 'this workspace'}
							</p>
						</div>
						<Button onclick={() => (addDialogOpen = true)}>
							<PlusIcon class="size-4" />
							Add Connection
						</Button>
					</div>
				</Card.Header>
				<Card.Content>
					{#if isLoading}
						<!-- Loading State -->
						<div class="grid gap-4 md:grid-cols-2 lg:grid-cols-3">
							{#each Array(3) as _}
								<div class="space-y-3 rounded-lg border p-4">
									<Skeleton class="h-5 w-2/3" />
									<Skeleton class="h-4 w-full" />
									<Skeleton class="h-4 w-1/2" />
								</div>
							{/each}
						</div>
					{:else if connectionList.length === 0}
						<!-- Empty State -->
						<div class="text-muted-foreground flex flex-col items-center justify-center py-12">
							<DatabaseIcon class="size-12 opacity-50" />
							<p class="mt-4 text-lg font-medium">No connections yet</p>
							<p class="mt-1 text-sm">Add your first data source connection to get started</p>
							<Button class="mt-4" onclick={() => (addDialogOpen = true)}>
								<PlusIcon class="size-4" />
								Add Connection
							</Button>
						</div>
					{:else}
						<!-- Connection Cards Grid -->
						<div class="grid gap-4 md:grid-cols-2 lg:grid-cols-3">
							{#each connectionList as connection (connection.id)}
								{@const status = getStatusBadge(getTestResult(connection))}
								{@const StatusIcon = status.icon}
								<Card.Root class="relative">
									<Card.Header class="pb-3">
										<div class="flex items-start justify-between">
											<Card.Title class="text-base">{connection.name}</Card.Title>
											<Badge variant={status.variant} class="text-xs">
												<StatusIcon class="size-3" />
												{status.label}
											</Badge>
										</div>
									</Card.Header>
									<Card.Content class="space-y-2">
										<div class="flex items-center gap-2">
											<Badge variant="secondary" class="text-xs">
												{connection.type.toUpperCase()}
											</Badge>
										</div>
										<p class="text-muted-foreground truncate text-sm">
											{getConnectionUrl(connection)}
										</p>
										<div class="pt-2">
											<Button
												variant="outline"
												size="sm"
												onclick={() => handleTest(connection.id)}
											>
												Test Connection
											</Button>
										</div>
									</Card.Content>
								</Card.Root>
							{/each}
						</div>
					{/if}
				</Card.Content>
			</Card.Root>
		</div>
	</Sidebar.Inset>
</Sidebar.Provider>

<!-- Add Connection Dialog -->
<Dialog.Root bind:open={addDialogOpen}>
	<Dialog.Content class="max-w-md">
		<Dialog.Header>
			<Dialog.Title>Add Connection</Dialog.Title>
			<Dialog.Description>
				Configure a new data source connection
			</Dialog.Description>
		</Dialog.Header>
		<form
			onsubmit={(e) => {
				e.preventDefault();
				handleSubmit();
			}}
		>
			<div class="space-y-4 py-4">
				<div class="space-y-2">
					<Label for="connection-name">Name *</Label>
					<Input
						id="connection-name"
						bind:value={formData.name}
						placeholder="My API Connection"
						disabled={isSubmitting}
					/>
				</div>

				<div class="space-y-2">
					<Label for="connection-type">Type *</Label>
					<Select.Root type="single" bind:value={formData.type} name="type">
						<Select.Trigger class="w-full">
							{connectionTypes.find(t => t.value === formData.type)?.label ?? 'Select type'}
						</Select.Trigger>
						<Select.Content>
							{#each connectionTypes as type}
								<Select.Item value={type.value} label={type.label}>
									{type.label}
								</Select.Item>
							{/each}
						</Select.Content>
					</Select.Root>
				</div>

				<div class="space-y-2">
					<Label for="connection-url">URL *</Label>
					<Input
						id="connection-url"
						bind:value={formData.url}
						placeholder="https://api.example.com"
						disabled={isSubmitting}
					/>
				</div>

				{#if formData.type === 'api'}
					<div class="space-y-2">
						<Label for="api-key">API Key</Label>
						<Input
							id="api-key"
							type="password"
							bind:value={formData.apiKey}
							placeholder="Enter API key"
							disabled={isSubmitting}
						/>
					</div>
				{:else}
					<div class="space-y-2">
						<Label for="db-username">Username</Label>
						<Input
							id="db-username"
							bind:value={formData.username}
							placeholder="Database username"
							disabled={isSubmitting}
						/>
					</div>

					<div class="space-y-2">
						<Label for="db-password">Password</Label>
						<Input
							id="db-password"
							type="password"
							bind:value={formData.password}
							placeholder="Database password"
							disabled={isSubmitting}
						/>
					</div>

					<div class="space-y-2">
						<Label for="db-name">Database Name</Label>
						<Input
							id="db-name"
							bind:value={formData.databaseName}
							placeholder="my_database"
							disabled={isSubmitting}
						/>
					</div>
				{/if}
			</div>

			<Dialog.Footer>
				<Button
					type="button"
					variant="outline"
					onclick={() => {
						addDialogOpen = false;
						resetForm();
					}}
					disabled={isSubmitting}
				>
					Cancel
				</Button>
				<Button
					type="button"
					variant="secondary"
					onclick={handleTestConnection}
					disabled={isTesting || isSubmitting}
				>
					{#if isTesting}
						<Loader2Icon class="size-4 animate-spin" />
					{/if}
					Test
				</Button>
				<Button type="submit" disabled={isSubmitting}>
					{#if isSubmitting}
						<Loader2Icon class="size-4 animate-spin" />
					{/if}
					Save
				</Button>
			</Dialog.Footer>
		</form>
	</Dialog.Content>
</Dialog.Root>
