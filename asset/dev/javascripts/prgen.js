function initPrGen() {
    document.getElementById("genpr").addEventListener("click", submit);
    document.getElementById("clear").addEventListener("click", clear);
    document.getElementById("copy").addEventListener("click", copyToClipBoard);
}

function submit() {
    let object = {
        Jira: document.getElementById("jira").value,
        LinkedPr: document.getElementById("linkedPr").value,
        RelatedPr: document.getElementById("relatedPr").value,
        Cms: document.getElementById('CMS').checked,
        Db: document.getElementById('DB').checked,
        LiveFollowUp: document.getElementById('LiveFollowUp').checked,
        Parameters: document.getElementById('Parameters').checked,
        Cloudinary: document.getElementById('Cloudinary').checked,
    };

    let xhr = new XMLHttpRequest();
    let url = "/";
    xhr.open("POST", url, true);
    xhr.setRequestHeader("Content-Type", "application/json");
    xhr.onreadystatechange = function () {
        if (xhr.readyState === 4 && xhr.status === 200) {
            document.getElementById('output').textContent = xhr.responseText;
        }
    }
    xhr.send(JSON.stringify(object));
}

function clear() {
    document.getElementById("jira").value = "";
    document.getElementById("linkedPr").value = "";
    document.getElementById("relatedPr").value = "";
    document.getElementById('CMS').checked = false;
    document.getElementById('DB').checked = false;
    document.getElementById('LiveFollowUp').checked = false;
    document.getElementById('Parameters').checked = false;
    document.getElementById('Cloudinary').checked = false;
}

function copyToClipBoard() {
    document.getElementById("output").select();
    document.execCommand("copy");

    let button = document.getElementById("copy");

    button.classList.remove("copy-to-clipboard-default");
    button.classList.add("copy-to-clipboard-success");

    setTimeout(function() {
        button.classList.remove("copy-to-clipboard-success");
        button.classList.add("copy-to-clipboard-default");
    }, 3000);
}

initPrGen();