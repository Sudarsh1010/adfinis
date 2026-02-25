import adapter from "@sveltejs/adapter-static";

/** @type {import('@sveltejs/kit').Config} */
const config = {
	kit: {
		adapter: adapter({ handleMissingId: "ignore" }),
		prerender: { handleMissingId: "ignore" },
		alias: {
			$components: "src/components",
			"$components/*": "src/components/*",
		},
	},
};

export default config;
