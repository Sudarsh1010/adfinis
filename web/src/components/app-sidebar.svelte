<script lang="ts" module>
	import type { Component } from "@lucide/svelte";
	import type { Pathname } from "$app/types";

	interface Data {
		user: {
			name: string;
			email: string;
			avatar: string;
		};
		navMain: (
			| {
					title: string;
					url: Pathname;
					icon: typeof Component;
					isActive: boolean;
					items: {
						title: string;
						url: Pathname;
					}[];
			  }
			| {
					title: string;
					url: Pathname;
					icon: typeof Component;
					items: {
						title: string;
						url: Pathname;
					}[];
					isActive?: undefined;
			  }
		)[];
		projects: {
			name: string;
			url: Pathname;
			icon: typeof Component;
		}[];
	}

	// This is sample data.
	const data: Data = {
		user: {
			name: "shadcn",
			email: "m@example.com",
			avatar: "/avatars/shadcn.jpg",
		},

		navMain: [
			// {
			// 	title: "Playground",
			// 	url: "#",
			// 	icon: SquareTerminalIcon,
			// 	isActive: true,
			// 	items: [
			// 		{
			// 			title: "History",
			// 			url: "#",
			// 		},
			// 		{
			// 			title: "Starred",
			// 			url: "#",
			// 		},
			// 		{
			// 			title: "Settings",
			// 			url: "#",
			// 		},
			// 	],
			// },
			// {
			// 	title: "Models",
			// 	url: "#",
			// 	icon: BotIcon,
			// 	items: [
			// 		{
			// 			title: "Genesis",
			// 			url: "#",
			// 		},
			// 		{
			// 			title: "Explorer",
			// 			url: "#",
			// 		},
			// 		{
			// 			title: "Quantum",
			// 			url: "#",
			// 		},
			// 	],
			// },
			// {
			// 	title: "Documentation",
			// 	url: "#",
			// 	icon: BookOpenIcon,
			// 	items: [
			// 		{
			// 			title: "Introduction",
			// 			url: "#",
			// 		},
			// 		{
			// 			title: "Get Started",
			// 			url: "#",
			// 		},
			// 		{
			// 			title: "Tutorials",
			// 			url: "#",
			// 		},
			// 		{
			// 			title: "Changelog",
			// 			url: "#",
			// 		},
			// 	],
			// },
			// {
			// 	title: "Settings",
			// 	url: "#",
			// 	icon: Settings2Icon,
			// 	items: [
			// 		{
			// 			title: "General",
			// 			url: "#",
			// 		},
			// 		{
			// 			title: "Team",
			// 			url: "#",
			// 		},
			// 		{
			// 			title: "Billing",
			// 			url: "#",
			// 		},
			// 		{
			// 			title: "Limits",
			// 			url: "#",
			// 		},
			// 	],
			// },
		],

		projects: [
			// {
			// 	name: "Design Engineering",
			// 	url: "#",
			// 	icon: FrameIcon,
			// },
			// {
			// 	name: "Sales & Marketing",
			// 	url: "#",
			// 	icon: ChartPieIcon,
			// },
			// {
			// 	name: "Travel",
			// 	url: "#",
			// 	icon: MapIcon,
			// },
		],
	};
</script>

<script lang="ts">
	import NavMain from "./nav-main.svelte";
	import NavProjects from "./nav-projects.svelte";
	import * as Sidebar from "$components/ui/sidebar/index.js";
	import type { ComponentProps } from "svelte";

	let {
		ref = $bindable(null),
		collapsible = "icon",
		...restProps
	}: ComponentProps<typeof Sidebar.Root> = $props();
</script>

<Sidebar.Root {collapsible} {...restProps}>
	<Sidebar.Content>
		<NavMain items={data.navMain} />
		<NavProjects projects={data.projects} />
	</Sidebar.Content>
	<Sidebar.Rail />
</Sidebar.Root>
