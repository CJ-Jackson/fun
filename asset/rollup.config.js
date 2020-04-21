import nodeResolve from 'rollup-plugin-node-resolve';
import commonjs from 'rollup-plugin-commonjs';
import { terser } from 'rollup-plugin-terser';

export default {
    input: {
        'main': 'dev/javascripts/main.js',
        'serviceWorker': 'dev/javascripts/serviceWorker.js',
    },
    output: {
        dir: 'live/javascripts',
        format: 'es',
        globals: [],
    },
    plugins: [
        nodeResolve(),
        commonjs(),
        terser()
    ],
    external: []
}