import Vue from 'vue';
import App from './App.vue'

if ("serviceWorker" in navigator) {
    window.addEventListener("load", function() {
        navigator.serviceWorker
            .register("/javascript/serviceWorker.js", {
                scope: '/'
            })
            .then(function (res) {
                return console.log("service worker registered");
            })
            .catch(function (err) {
                return console.log("service worker not registered", err);
            })
    });
}

let application = Vue.extend(App);
let vm = new application({
    el: "#app"
});