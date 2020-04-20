import './global';
import 'bootstrap';
import './prgen';

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