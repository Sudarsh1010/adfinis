<script lang="ts">
	import { workspaceStore } from '$lib/stores/workspace';
	import {
		Select,
		SelectContent,
		SelectItem,
		SelectTrigger
	} from '$components/ui/select';

	// Fetch workspaces on mount
	$effect(() => {
		workspaceStore.fetchWorkspaces();
	});

	function handleValueChange(workspaceId: string) {
		workspaceStore.selectWorkspace(workspaceId);
	}
</script>

<Select
	type="single"
	value={workspaceStore.currentWorkspace?.id ?? ''}
	onValueChange={handleValueChange}
>
	<SelectTrigger>
		{workspaceStore.currentWorkspace?.name ?? 'Select workspace'}
	</SelectTrigger>
	<SelectContent>
		{#if workspaceStore.workspaces.length === 0}
			<SelectItem value="__none" label="No workspaces">
				<span class="text-muted-foreground">No workspaces</span>
				<a href="/workspaces" class="text-primary ml-2 underline">Create one</a>
			</SelectItem>
		{:else}
			{#each workspaceStore.workspaces as workspace (workspace.id)}
				<SelectItem value={workspace.id} label={workspace.name}>
					{workspace.name}
				</SelectItem>
			{/each}
		{/if}
	</SelectContent>
</Select>
