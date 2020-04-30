import './global';
import 'bootstrap';

import Vue from 'vue';
import App from './App.vue'

if ("serviceWorker" in navigator) {
    window.addEventListener("load", function() {
        navigator.serviceWorker
            .register("/javascripts/serviceWorker.js", {
                scope: '/'
            })
            .then(res => console.log("service worker registered"))
            .catch(err => console.log("service worker not registered", err))
    });
}

let application = Vue.extend(App);
let vm = new application({
    el: "#app"
});