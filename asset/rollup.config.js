import nodeResolve from 'rollup-plugin-node-resolve';
import commonjs from 'rollup-plugin-commonjs';

export default {
    input: {
        'main': 'dev/javascripts/main.js',
    },
    output: {
        dir: 'live/javascripts',
        format: 'es',
        globals: [],
    },
    plugins: [
        nodeResolve(),
        commonjs()
    ],
    external: []
}