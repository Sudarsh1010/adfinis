<script lang="ts">
	import { z } from "zod";
	import { Button } from "$components/ui/button/index.js";
	import { Input } from "$components/ui/input/index.js";
	import { Badge } from "$components/ui/badge/index.js";
	import {
		Select,
		SelectContent,
		SelectItem,
		SelectTrigger
	} from "$components/ui/select/index.js";
	import { Label } from "$components/ui/label/index.js";

	// Connection status type
	export type ConnectionStatus = "idle" | "testing" | "connected" | "error";

	// Form data interface
	export interface ConnectionFormData {
		name: string;
		provider: "milvus";
		endpoint: string;
		apiKey?: string;
	}

	// Props interface
	export interface ConnectionFormProps {
		onSubmit: (data: ConnectionFormData) => void | Promise<void>;
		onTestConnection: (data: ConnectionFormData) => Promise<boolean>;
		initialData?: Partial<ConnectionFormData>;
		status?: ConnectionStatus;
	}

	let {
		onSubmit,
		onTestConnection,
		initialData = {},
		status = "idle"
	}: ConnectionFormProps = $props();

	// Zod schema for validation
	const schema = z.object({
		name: z.string().min(1, "Name is required"),
		provider: z.literal("milvus"),
		endpoint: z.string().url("Must be a valid URL"),
		apiKey: z.string().optional()
	});

	// Form state
	let formData = $state<ConnectionFormData>({
		name: initialData.name ?? "",
		provider: initialData.provider ?? "milvus",
		endpoint: initialData.endpoint ?? "",
		apiKey: initialData.apiKey ?? ""
	});

	// Validation errors
	let errors = $state<Partial<Record<keyof ConnectionFormData, string>>>({});

	// Submission state
	let isSubmitting = $state(false);

	// Validate a single field
	function validateField(field: keyof ConnectionFormData): boolean {
		const result = schema.shape[field].safeParse(formData[field]);
		if (!result.success) {
			errors[field] = result.error.issues[0]?.message;
			return false;
		}
		errors[field] = undefined;
		return true;
	}

	// Validate all fields
	function validateForm(): boolean {
		const result = schema.safeParse(formData);
		if (!result.success) {
			errors = {};
			for (const issue of result.error.issues) {
				const field = issue.path[0] as keyof ConnectionFormData;
				errors[field] = issue.message;
			}
			return false;
		}
		errors = {};
		return true;
	}

	// Handle form submission
	async function handleSubmit(e: Event) {
		e.preventDefault();
		if (!validateForm()) return;

		isSubmitting = true;
		try {
			await onSubmit(formData);
		} finally {
			isSubmitting = false;
		}
	}

	// Handle test connection
	async function handleTestConnection() {
		// Only validate name and endpoint for test
		const testSchema = z.object({
			name: schema.shape.name,
			endpoint: schema.shape.endpoint
		});
		
		const result = testSchema.safeParse({ 
			name: formData.name, 
			endpoint: formData.endpoint 
		});
		
		if (!result.success) {
			errors = {};
			for (const issue of result.error.issues) {
				const field = issue.path[0] as keyof ConnectionFormData;
				errors[field] = issue.message;
			}
			return;
		}
		errors = {};

		await onTestConnection(formData);
	}

	// Status badge configuration
	function getStatusBadgeVariant(): "default" | "secondary" | "destructive" | "outline" {
		switch (status) {
			case "connected":
				return "default";
			case "error":
				return "destructive";
			case "testing":
				return "secondary";
			default:
				return "outline";
		}
	}

	function getStatusLabel(): string {
		switch (status) {
			case "connected":
				return "Connected";
			case "error":
				return "Error";
			case "testing":
				return "Testing...";
			default:
				return "Not Connected";
		}
	}
</script>

<form onsubmit={handleSubmit} class="space-y-4">
	<!-- Status Badge -->
	<div class="flex items-center justify-between">
		<Label class="text-base font-semibold">Connection Details</Label>
		<Badge variant={getStatusBadgeVariant()}>{getStatusLabel()}</Badge>
	</div>

	<!-- Name Field -->
	<div class="space-y-2">
		<Label for="name">Name</Label>
		<Input
			id="name"
			type="text"
			placeholder="My Connection"
			bind:value={formData.name}
			onblur={() => validateField("name")}
			aria-invalid={!!errors.name}
		/>
		{#if errors.name}
			<p class="text-destructive text-sm font-medium">{errors.name}</p>
		{/if}
	</div>

	<!-- Provider Field -->
	<div class="space-y-2">
		<Label for="provider">Provider</Label>
		<Select
			type="single"
			value={formData.provider}
			onValueChange={(value: string) => {
				formData.provider = value as "milvus";
			}}
		>
			<SelectTrigger id="provider" class="w-full">
				{#if formData.provider === "milvus"}
					Milvus
				{:else}
					Select provider
				{/if}
			</SelectTrigger>
			<SelectContent>
				<SelectItem value="milvus" label="Milvus">
					Milvus
				</SelectItem>
			</SelectContent>
		</Select>
	</div>

	<!-- Endpoint Field -->
	<div class="space-y-2">
		<Label for="endpoint">Endpoint</Label>
		<Input
			id="endpoint"
			type="url"
			placeholder="https://localhost:19530"
			bind:value={formData.endpoint}
			onblur={() => validateField("endpoint")}
			aria-invalid={!!errors.endpoint}
		/>
		{#if errors.endpoint}
			<p class="text-destructive text-sm font-medium">{errors.endpoint}</p>
		{/if}
	</div>

	<!-- API Key Field -->
	<div class="space-y-2">
		<Label for="apiKey">API Key (Optional)</Label>
		<Input
			id="apiKey"
			type="password"
			placeholder="Enter your API key"
			bind:value={formData.apiKey}
		/>
	</div>

	<!-- Action Buttons -->
	<div class="flex gap-2 pt-2">
		<Button type="button" variant="outline" onclick={handleTestConnection} disabled={status === "testing"}>
			{#if status === "testing"}
				Testing...
			{:else}
				Test Connection
			{/if}
		</Button>
		<Button type="submit" disabled={isSubmitting}>
			{#if isSubmitting}
				Saving...
			{:else}
				Save Connection
			{/if}
		</Button>
	</div>
</form>
