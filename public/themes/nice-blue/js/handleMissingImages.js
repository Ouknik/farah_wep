
var missingImageElements = {};
var handleMissingImages = function(element) {
    // console.log(element.dataset.id)
    missingImageElements[element.dataset.id] = element;
    prepareShowBusy(element);
}

window.onload = function () {
    fetchMissingImages();
};

function fetchMissingImages() {
    if (Object.keys(missingImageElements).length == 0) return;
    const formData = [];

    Object.keys(missingImageElements).forEach(function(i) {
        formData.push({"id" : missingImageElements[i].dataset.id, "name": missingImageElements[i].dataset.name, "type" : missingImageElements[i].dataset.type});
    });

    var xmlhttp = new XMLHttpRequest();   // new HttpRequest instance 
    xmlhttp.open("POST", "/commonwebservices/azuredownloadmissingimages");
    xmlhttp.setRequestHeader("Content-Type", "application/json");
    xmlhttp.send(JSON.stringify(formData));
    xmlhttp.onreadystatechange = function() {
        // Check if the request is compete and was successful
        if (this.readyState === 4) {
            handleSuccessDownloadedImages(this);
        }
    };


}

function handleSuccessDownloadedImages(responseData) {
    // console.log(responseData.response);
    const response = JSON.parse(responseData.response);
    // console.log(response);
    if (response["success"] == 1) {
        const imageData = JSON.parse(response["data"]);
        // console.log(imageData)
        Object.keys(imageData).forEach(function(i, key) {
            // console.log(i)
            (imageData[i] != 1) ? appendDownloadMissingImageError(missingImageElements[i]) : showImage(missingImageElements[i]);
        });
    }
}


function appendDownloadMissingImageError(imageElement) {
    errorElement = document.createElement("p");
    errorElement.innerHTML = "Das Bild konnte nicht geladen werden. Möglicherweise wurde das Bild von Azure gelöscht.";
    errorElement.classList.add("missing-image");
    // use the document rather imageelement as imageeelement may not be loaded
    document.getElementsByClassName("downloading-image-progress")[0].remove();
    imageElement.parentElement.appendChild(errorElement);
    imageElement.remove();
}


function prepareShowBusy(imageElement) {
    showBuys = document.createElement("div");
    showBuys.innerHTML = "<div class='preloader-wrapper small active'><div class='spinner-layer _spinner-pink-only'><div class='circle-clipper left'><div class='circle'></div></div><div class='gap-patch'><div class='circle'></div></div><div class='circle-clipper right'><div class='circle'></div></div></div></div><div class='m-t-20' style='width: 25px; height: 25px;'>Das Bild wird heruntergeladen.</div>";
    showBuys.classList.add("downloading-image-progress");
    imageElement.parentElement.appendChild(showBuys)
    imageElement.classList.add("hide");
}


function showImage(imageElement) {
    // console.log(imageElement.parentElement.getElementsByClassName("downloading-image-progress")[0]);
    imageElement.src = imageElement.src
    imageElement.classList.remove("hide");
    imageElement.parentElement.getElementsByClassName("downloading-image-progress")[0].remove();
}