import nodeResolve from 'rollup-plugin-node-resolve';
import commonjs from 'rollup-plugin-commonjs';
import { terser } from 'rollup-plugin-terser';
import vue from 'rollup-plugin-vue';
import replace from '@rollup/plugin-replace';

export default {
    input: {
        'main': 'dev/javascript/main.js',
        'serviceWorker': 'dev/javascript/serviceWorker.js',
    },
    output: {
        dir: 'live/javascript',
        format: 'es',
        globals: [],
    },
    plugins: [
        vue(),
        nodeResolve(),
        commonjs(),
        terser(),
        replace({
            'process.env.NODE_ENV': JSON.stringify( process.env.BUILD )
        })
    ],
    external: []
}