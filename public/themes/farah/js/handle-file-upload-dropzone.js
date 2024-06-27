

var previewNode = document.querySelector("#dropzone-template");
previewNode.id = "";
var previewTemplate = previewNode.parentNode.innerHTML;
previewNode.parentNode.removeChild(previewNode);

// Disable auto discover for all elements:
Dropzone.autoDiscover = false;

var _dropzoneUploadFile = null;

// init the sweet alert
var _instanceSweetalert = new sweetalert();


function initDropzone() {

    var params = {};

    var _acceptedFileTypes = "image/*";

    var _successCallback = handleSuccess;

    if (typeof acceptedFileTypes !== "undefined") {
        _acceptedFileTypes = acceptedFileTypes;
    }

    if (typeof formDataToSend !== "undefined") {
        params = formDataToSend;
    }

    if (typeof successCallback !== "undefined") {
        _successCallback = successCallback
    }


    _dropzoneUploadFile = new Dropzone("#fileUploadDropzone", {

        url: dropzoneFileUploadUrl,
        paramName: dropzoneFileUploadParamName,

        //createImageThumbnails: false,
        //parallelUploads: 1,
        uploadMultiple: false,
        clickable: ".dropzone-trigger",

        previewTemplate: previewTemplate,
        previewsContainer: "#previews", // Define the container to display the previews

        maxFilesize: 10,
        maxFiles: 1,
        acceptedFiles: _acceptedFileTypes,

        autoProcessQueue: true,

        params : params,


        init: function () {

            // create drop here overlay effect
            this.on('dragenter', function(file) { handleDragEnter(); });
            // hide drop here overlay
            this.on("dragleave", function(file) { handleDragOut(); });
            // keep it idle if you need or remove
            this.on("addedfile", function(file) { handleDragOut(); });
            // handle error
            // this.on("error", function(file) { console.log("error dear"); });

            this.on("thumbnail", function(file) {

                handleDragOut();

                if (file.status == "error") return;

                // this.processQueue();
            });

            this.on('success', function (file, response) {
                //console.log(response);
                _instanceSweetalert.close();
                response = JSON.parse(response);
                setTimeout(
                    function() {
                        // console.log(response);
                        if (response["error"] == "1") {
                            // console.log(file.error);
                            _instanceSweetalert.init("Oops!", response["message"], true, "Ok", true);
                            _instanceSweetalert.setType("error");
                        } else {
                            // handleSuccess(response);
                            _successCallback(response);
                            _instanceSweetalert.init("Fertig!", response["message"]);
                            _instanceSweetalert.setType("success");
                            _instanceSweetalert.close(2000);
                        }
                        _instanceSweetalert.show();
                    }, 1000
                );
            });

            this.on("totaluploadprogress", function(progress) {
                //console.log(progress + "%");
                _instanceSweetalert.updateProgressbar(progress.toFixed(2));
            });

            this.on("sending", function(file, xhr, formData){
                // set company id
                triggerAvatarUploadProgressModal();
            });

            this.on("complete", function(file) {
                //console.log("complete");
                if (file.status == "error") {
                    document.querySelector(".dropzone-container-error").classList.remove("hide");
                } else {
                    //_dropzoneUploadFile.removeAllFiles();
                    //_dropzoneUploadFile.removeAllFiles(true);
                    destroy();
                }
            });
            /*
            this.on("error", function(error, errorMessage) {
            console.log(error);
            console.log(errorMessage);
            });
            */
        }
    });
}


function handleDragEnter() {
    document.querySelector(".dropzone-container-dragover").classList.remove("hide");
    document.querySelector(".dropzone-container-dragover-content").classList.add("slide-down");
    // hide error if displayed
    document.querySelector(".dropzone-container-error").classList.add("hide");
}

function handleDragOut() {
    document.querySelector(".dropzone-container-dragover").classList.add("hide");
    document.querySelector(".dropzone-container-dragover-content").classList.remove("slide-down");
}

function triggerAvatarUploadProgressModal() {
    _instanceSweetalert.showProgressbar("Einen Moment bitte...");
}




function destroy() {
    //console.log("destroy");
    // _dropzoneUploadFile.disable();
    // _dropzoneUploadFile.enable();
    _dropzoneUploadFile.removeAllFiles();
    _dropzoneUploadFile.removeAllFiles(true);
    _dropzoneUploadFile.destroy();
    _dropzoneUploadFile = null;
    initDropzone();
}


function handleSuccess(response) {
    try {
        // there may have had been an error before upload
        // therefore hide the error message
        document.querySelector(".dropzone-container-error").classList.add("hide");
        // remove the dropzone icon (background)
        document.querySelector(".dropzone-container-background").classList.add("hide");
        // Unhide the delete avatar trigger button
        document.getElementById("trigger-delete-uploaded-image").classList.remove("hide");

        // set the profile pic
        setTimeout(function(){
            document.querySelector(".dropzone-cotainer-file-upload").innerHTML = `<img src="${response.image}">`;
        },500);

    } catch (e) {
        console.log(e);
    }
}

initDropzone();





//---------------------------------------------------------------//
//---------------------BEGIN DELETE PROJECT IMAGE-----------------------//
document.getElementById("trigger-delete-uploaded-image").onclick = function(e) {
    e.preventDefault();
    this.blur();
    // console.log("Bismillah" + dropzoneFileDeleteUrl);
    _instanceSweetalert.showBusy("Bitte warten...");
    setTimeout(
        function() {
            ajax(dropzoneFileDeleteUrl, "POST", null, deleteHandleSuccess);
        },500
    );
};

deleteHandleSuccess = function(data) {
    var response = JSON.parse(data.responseText);
    if (response.success == 1) {
        try {
            // hide the delete trigger
            document.getElementById("trigger-delete-uploaded-image").classList.add("hide");
            // unhide the camera icon in background
            document.querySelector(".dropzone-container-background").classList.remove("hide");
            // remove the image from dropzone container
            document.querySelector(".dropzone-cotainer-file-upload").innerHTML = "";

            setTimeout(function() {
                //_instanceSweetalert.close();
                _instanceSweetalert.init("", "", true);
                _instanceSweetalert.setType("success");
                _instanceSweetalert.show();
                setTimeout(
                    function() {
                        _instanceSweetalert.close();
                    }, 2000
                )
            }, 500);
        } catch(e) {
            console.log(e)
        }
    }
}
//---------------------ENd DELETE PROJECT IMAGE-----------------------//
//---------------------------------------------------------------//
