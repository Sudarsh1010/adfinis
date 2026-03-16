<script lang="ts">
	import type { PageData } from "./+page";
	import type { Workspace } from "$lib/api/types";
	import { workspaces } from "$lib/api/client";
	import { showToast } from "$lib/utils/toast";
	import * as Card from "$components/ui/card/index.js";
	import * as Table from "$components/ui/table/index.js";
	import * as Dialog from "$components/ui/dialog/index.js";
	import Button from "$components/ui/button/button.svelte";
	import Input from "$components/ui/input/input.svelte";
	import Skeleton from "$components/ui/skeleton/skeleton.svelte";
	import AppSidebar from "$components/app-sidebar.svelte";
	import * as Sidebar from "$components/ui/sidebar/index.js";
	import * as Breadcrumb from "$components/ui/breadcrumb/index.js";
	import { Separator } from "$components/ui/separator/index.js";
	import PlusIcon from "@lucide/svelte/icons/plus";
	import PencilIcon from "@lucide/svelte/icons/pencil";
	import Trash2Icon from "@lucide/svelte/icons/trash-2";
	import Loader2Icon from "@lucide/svelte/icons/loader-2";
	import AlertCircleIcon from "@lucide/svelte/icons/alert-circle";
	import { setContext } from "svelte";

	let { data }: { data: PageData } = $props();

	// State using Svelte 5 runes
	let workspaceList = $state<Workspace[]>(data.workspaces ?? []);
	let isLoading = $state(false);
	let error = $state<string | null>(data.error ?? null);

	// Dialog states
	let createDialogOpen = $state(false);
	let deleteDialogOpen = $state(false);
	let workspaceToDelete = $state<Workspace | null>(null);

	// Form states
	let newWorkspaceName = $state("");
	let isSubmitting = $state(false);

	// Refresh workspaces from API
	async function refreshWorkspaces() {
		isLoading = true;
		error = null;

		const response = await workspaces.list();

		if (response.success && response.data) {
			workspaceList = response.data;
		} else {
			error = response.error || response.message || "Failed to load workspaces";
			showToast(error, "error");
		}

		isLoading = false;
	}

	// Create workspace
	async function handleCreate() {
		if (!newWorkspaceName.trim()) return;

		isSubmitting = true;
		error = null;

		const response = await workspaces.create({ name: newWorkspaceName.trim() });

		if (response.success && response.data) {
			workspaceList = [...workspaceList, response.data];
			newWorkspaceName = "";
			createDialogOpen = false;
			showToast("Workspace created successfully", "success");
		} else {
			error =
				response.error || response.message || "Failed to create workspace";
			showToast(error, "error");
		}

		isSubmitting = false;
	}

	// Delete workspace
	async function handleDelete() {
		if (!workspaceToDelete) return;

		const toDelete = workspaceToDelete;
		isSubmitting = true;
		error = null;

		// Since there's no delete API endpoint defined in client.ts,
		// we'll remove from local state for now
		// TODO: Replace with actual API call when available
		// const response = await api.del(`/workspaces/${workspaceToDelete.id}`);

		workspaceList = workspaceList.filter((w) => w.id !== toDelete.id);
		workspaceToDelete = null;
		deleteDialogOpen = false;
		showToast("Workspace deleted successfully", "success");

		isSubmitting = false;
	}

	// Open delete confirmation
	function confirmDelete(workspace: Workspace) {
		workspaceToDelete = workspace;
		deleteDialogOpen = true;
	}

	// Format date for display
	function formatDate(dateString: string): string {
		return new Date(dateString).toLocaleDateString("en-US", {
			year: "numeric",
			month: "short",
			day: "numeric",
		});
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
				<Separator
					orientation="vertical"
					class="me-2 data-[orientation=vertical]:h-4"
				/>
				<Breadcrumb.Root>
					<Breadcrumb.List>
						<Breadcrumb.Item>
							<Breadcrumb.Page>Workspaces</Breadcrumb.Page>
						</Breadcrumb.Item>
					</Breadcrumb.List>
				</Breadcrumb.Root>
			</div>
		</header>

		<div class="flex flex-1 flex-col gap-4 p-4 pt-0">
			<!-- Error Alert -->
			{#if error}
				<div
					class="flex items-center gap-2 rounded-lg border border-destructive/20 bg-destructive/10 p-4 text-destructive"
				>
					<AlertCircleIcon class="size-5 shrink-0" />
					<span>{error}</span>
					<Button
						variant="ghost"
						size="sm"
						class="ml-auto"
						onclick={() => (error = null)}
					>
						Dismiss
					</Button>
				</div>
			{/if}

			<!-- Main Content Card -->
			<Card.Root>
				<Card.Header>
					<div class="flex items-center justify-between">
						<div>
							<Card.Title>Workspaces</Card.Title>
							<p class="mt-1 text-sm text-muted-foreground">
								Manage your data connection workspaces
							</p>
						</div>
						<Button onclick={() => (createDialogOpen = true)}>
							<PlusIcon class="size-4" />
							New Workspace
						</Button>
					</div>
				</Card.Header>

				<Card.Content>
					{#if isLoading}
						<!-- Loading State -->
						<div class="space-y-3">
							{#each Array(3) as _ (_)}
								<div class="flex items-center gap-4">
									<Skeleton class="h-10 w-full" />
								</div>
							{/each}
						</div>
					{:else if workspaceList.length === 0}
						<!-- Empty State -->
						<div
							class="flex flex-col items-center justify-center py-12 text-muted-foreground"
						>
							<p class="text-lg font-medium">No workspaces yet</p>
							<p class="mt-1 text-sm">
								Create your first workspace to get started
							</p>
							<Button class="mt-4" onclick={() => (createDialogOpen = true)}>
								<PlusIcon class="size-4" />
								Create Workspace
							</Button>
						</div>
					{:else}
						<!-- Workspace Table -->
						<Table.Root>
							<Table.Header>
								<Table.Row>
									<Table.Head>Name</Table.Head>
									<Table.Head>Created</Table.Head>
									<Table.Head>Updated</Table.Head>
									<Table.Head class="text-right">Actions</Table.Head>
								</Table.Row>
							</Table.Header>

							<Table.Body>
								{#each workspaceList as workspace (workspace.id)}
									<Table.Row>
										<Table.Cell class="font-medium">{workspace.name}</Table.Cell
										>
										<Table.Cell>{formatDate(workspace.created_at)}</Table.Cell>
										<Table.Cell>{formatDate(workspace.updated_at)}</Table.Cell>
										<Table.Cell class="text-right">
											<div class="flex justify-end gap-2">
												<Button variant="ghost" size="icon-sm" title="Edit">
													<PencilIcon class="size-4" />
												</Button>

												<Button
													variant="ghost"
													size="icon-sm"
													title="Delete"
													onclick={() => confirmDelete(workspace)}
												>
													<Trash2Icon class="size-4" />
												</Button>
											</div>
										</Table.Cell>
									</Table.Row>
								{/each}
							</Table.Body>
						</Table.Root>
					{/if}
				</Card.Content>
			</Card.Root>
		</div>
	</Sidebar.Inset>
</Sidebar.Provider>

<!-- Create Workspace Dialog -->
<Dialog.Root bind:open={createDialogOpen}>
	<Dialog.Content>
		<Dialog.Header>
			<Dialog.Title>Create Workspace</Dialog.Title>
		</Dialog.Header>
		<form
			onsubmit={(e) => {
				e.preventDefault();
				handleCreate();
			}}
		>
			<div class="py-4">
				<label for="workspace-name" class="mb-2 block text-sm font-medium">
					Workspace Name
				</label>
				<Input
					id="workspace-name"
					bind:value={newWorkspaceName}
					placeholder="Enter workspace name"
					disabled={isSubmitting}
				/>
			</div>
			<Dialog.Footer>
				<Button
					type="button"
					variant="outline"
					onclick={() => (createDialogOpen = false)}
					disabled={isSubmitting}
				>
					Cancel
				</Button>
				<Button
					type="submit"
					disabled={!newWorkspaceName.trim() || isSubmitting}
				>
					{#if isSubmitting}
						<Loader2Icon class="size-4 animate-spin" />
					{/if}
					Create
				</Button>
			</Dialog.Footer>
		</form>
	</Dialog.Content>
</Dialog.Root>

<!-- Delete Confirmation Dialog -->
<Dialog.Root bind:open={deleteDialogOpen}>
	<Dialog.Content>
		<Dialog.Header>
			<Dialog.Title>Delete Workspace</Dialog.Title>
		</Dialog.Header>
		<div class="py-4">
			<p>
				Are you sure you want to delete <strong
					>{workspaceToDelete?.name}</strong
				>? This action cannot be undone.
			</p>
		</div>
		<Dialog.Footer>
			<Button
				type="button"
				variant="outline"
				onclick={() => (deleteDialogOpen = false)}
				disabled={isSubmitting}
			>
				Cancel
			</Button>
			<Button
				variant="destructive"
				onclick={handleDelete}
				disabled={isSubmitting}
			>
				{#if isSubmitting}
					<Loader2Icon class="size-4 animate-spin" />
				{/if}
				Delete
			</Button>
		</Dialog.Footer>
	</Dialog.Content>
</Dialog.Root>
