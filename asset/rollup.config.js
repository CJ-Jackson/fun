import nodeResolve from 'rollup-plugin-node-resolve';
import commonjs from 'rollup-plugin-commonjs';
import { terser } from 'rollup-plugin-terser';
import vue from 'rollup-plugin-vue';
import replace from '@rollup/plugin-replace';

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
        vue(),
        nodeResolve(),
        commonjs(),
        terser(),
        replace({
            'process.env.NODE_ENV': JSON.stringify( 'production' )
        })
    ],
    external: []
}