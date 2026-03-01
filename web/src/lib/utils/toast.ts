import { toast } from "svelte-sonner";

export type ToastType = "success" | "error" | "info";

export function showToast(message: string, type: ToastType = "info") {
	switch (type) {
		case "success":
			toast.success(message);
			break;
		case "error":
			toast.error(message);
			break;
		case "info":
		default:
			toast.info(message);
			break;
	}
}
