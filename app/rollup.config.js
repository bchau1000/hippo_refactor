// rollup.config.js
import alias from '@rollup/plugin-alias';

const aliases = alias({
	resolve: ['.svelte', '.ts', '.js'], //optional, by default this will just look for .js files or folders
	entries: [
		{ find: 'components', replacement: 'src/components' },
		{ find: 'metadata', replacement: 'src/metadata' },
		{ find: 'util', replacement: 'src/util' }
	]
});

export default {
	plugins: [aliases]
};
