

var _instanceModalAvatarCropper;
var fileCropped = false;
var boundsToCrop;


var previewNode = document.querySelector("#dropzone-template");
previewNode.id = "";
var previewTemplate = previewNode.parentNode.innerHTML;
previewNode.parentNode.removeChild(previewNode);


// Disable auto discover for all elements:
Dropzone.autoDiscover = false;

var _dropzoneUploadAvatar = null;

// init the sweet alert
var _instanceSweetalert = new sweetalert();


function initDropzone() {
  _dropzoneUploadAvatar = new Dropzone("#avatarDropzone", {

    url: "/commonwebservices/uploadavatar",
    paramName: "Avatar",
    /*
    thumbnailWidth: 180,
    thumbnailHeight: 180,
    */
    //createImageThumbnails: false,
    //parallelUploads: 1,
    uploadMultiple: false,
    clickable: ".dropzone-trigger",

    previewTemplate: previewTemplate,
    previewsContainer: "#previews", // Define the container to display the previews

    maxFilesize: 10,
    maxFiles: 1,
    acceptedFiles: "image/*",

    autoProcessQueue: false,



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
      
        if (fileCropped) {
          fileCropped = false;
          this.processQueue();
          //console.log("Processed");
        } else {
          handleCrop(file);
        }
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
              handleSuccess(response);
              _instanceSweetalert.init("Fertig!", response["message"]);
              _instanceSweetalert.setType("success");
              _instanceSweetalert.close(4000);
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
        //console.log("Sending" + boundsToCrop);
        //console.log(boundsToCrop.rotate);
        // fix the cropper rotation direction bug
        if (boundsToCrop.rotate == -90 || boundsToCrop.rotate == 90) {
          boundsToCrop.rotate = boundsToCrop.rotate * -1;
        }
        formData.append('boundsToCrop', JSON.stringify(boundsToCrop));
        boundsToCrop = null;
      });

      this.on("complete", function(file) {
        //console.log("complete");
        if (file.status == "error") {
          document.querySelector(".dropzone-container-error").classList.remove("hide");
        } else {
          //_dropzoneUploadAvatar.removeAllFiles();
          //_dropzoneUploadAvatar.removeAllFiles(true);
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

function triggerCropperModal() {
  var modalAvatarContainer = document.querySelector("#modalAvatarCropper");
  _instanceModalAvatarCropper = M.Modal.init(modalAvatarContainer, {dismissible: false});
  _instanceModalAvatarCropper.open();
}

function triggerAvatarUploadProgressModal() {
  _instanceSweetalert.showProgressbar("Bitte warten...");
}

function closeAvatarCropperModal() {
  try {
    setTimeout(function() {
      _instanceModalAvatarCropper.close();
    }, 100);
  } catch(e) {console.log(e);}
}

/*
function closeAvatarUploadProgressModal() {
  try {
    setTimeout(function() {
      _instanceSweetalert.close();
    }, 100);
  } catch(e) {console.log(e);}
}
*/



function handleCrop(file) {

  //_dropzoneUploadAvatar.removeFile(file);

  // Init the settings
  var editor = document.querySelector("#cropperjsContainer");
  var buttonDone = document.querySelector("#cropperjsDone");
  var buttonCancel = document.querySelector("#cropperjsCancel");

  var buttonRotateLeft = document.getElementById("trigger-rotate-left-avatar");
  var buttonRotateRight = document.getElementById("trigger-rotate-right-avatar");


  // Init Modal for cropper
  triggerCropperModal();

  // Create an image node for Cropper.js
  var image = new Image();
  image.src = URL.createObjectURL(file);
  editor.appendChild(image);
  var cropBoxData;
  var canvasData;

  // Create Cropper.js
  cropper = new Cropper(image, { 
    zoomOnWheel: false,
    rotatable: true,
    background: false,
    aspectRatio: 1,

    ready: function () {
      //Should set crop box data first here
      cropper.setCropBoxData(cropBoxData).setCanvasData(canvasData);
      //console.log(cropper.getImageData());
      //cropper.zoom(-0.1);
    },
  });

  buttonCancel.onclick = function() {
    cancel(cropper);
  };

  buttonDone.onclick = function() {    // Get the canvas with image data from Cropper.js
    var canvas = cropper.getCroppedCanvas({
      /*
      width: 256,
      height: 256,
      // Quality adjustments
      // imageSmoothingEnabled: false,
      // imageSmoothingQuality: 'high',
      */
      minWidth: 256,
      minHeight: 256,
      maxWidth: 2048,
      maxHeight: 2048,
    });

/*
    // if you want to upload cropped image then uncomment this
    // block and comment the bottom block
    // --------
    // Turn the canvas into a Blob (file object without a name)
    canvas.toBlob(function(blob) {    // Return the file to Dropzone
      blob.name = file.name;
      _dropzoneUploadAvatar.addFile(blob);
      fileCropped = true;
      closeAvatarCropperModal();
      triggerAvatarUploadProgressModal();
      cropper.reset();
      document.querySelector("#cropperjsContainer").innerHTML = "";
    });
*/
    
    fileCropped = true;
//    boundsToCrop = JSON.stringify(cropper.getCropBoxData());
    //boundsToCrop = JSON.stringify(cropper.getData(true));
    boundsToCrop = cropper.getData(true);
    /*
    let croppedData = cropper.getCropBoxData();
    boundsToCrop = new Array();
    boundsToCrop["left"] = croppedData.left;
    boundsToCrop["top"] = croppedData.top;
    boundsToCrop["width"] = croppedData.width;
    boundsToCrop["height"] = croppedData.height;
    */
    _dropzoneUploadAvatar.processQueue(file);
    closeAvatarCropperModal();
    triggerAvatarUploadProgressModal();
    cropper.reset();
    document.querySelector("#cropperjsContainer").innerHTML = "";

  };

  buttonRotateLeft.onclick = function() {
    cropper.rotate(-90);
  }

  buttonRotateRight.onclick = function() {
    cropper.rotate(90);
  }
}

function cancel(cropper) {
  cropper.destroy();
  document.querySelector("#cropperjsContainer").innerHTML = "";
  destroy();
}



function destroy() {
  //console.log("destroy");
  // _dropzoneUploadAvatar.disable();
  // _dropzoneUploadAvatar.enable();
  fileCropped = false;
  _dropzoneUploadAvatar.removeAllFiles();
  _dropzoneUploadAvatar.removeAllFiles(true);
  _dropzoneUploadAvatar.destroy();
  _dropzoneUploadAvatar = null;
  initDropzone();
  closeAvatarCropperModal();
}


function handleSuccess(response) {
  try {
    // there may have had been an error before upload
    // therefore hide the error message
    document.querySelector(".dropzone-container-error").classList.add("hide");
    // remove the dropzone icon (background)
    document.querySelector(".dropzone-container-background").classList.add("hide");
    // Unhide the delete avatar trigger button
    document.getElementById("trigger-delete-avatar").classList.remove("hide");

    // set the profile pic
    setTimeout(function(){
      document.querySelector(".dropzone-cotainer-profilepic").innerHTML = `<img src="${response.avatar}">`;


      document.querySelector(".top-nav-user-menu-trigger").innerHTML = '<img src="' + response.thumbnail + '" style="width: 25px; height: 25px; vertical-align: middle;">';


    },500);
    
  } catch (e) {
    //console.log(e);
  }
}

initDropzone();





//---------------------------------------------------------------//
//---------------------BEGIN DELETE AVATAR-----------------------//
document.getElementById("trigger-delete-avatar").onclick = function(e) {
    e.preventDefault();
    this.blur();
    _instanceSweetalert.showBusy("Bitte warten...");
    setTimeout(
      function() {
        ajax("/commonwebservices/deleteavatar", "POST", null, deleteAvatarHandleSuccess);
      },500
    );
    //ajax("/commonwebservices/deleteavatar", "POST", null, deleteAvatarHandleSuccess);
};

deleteAvatarHandleSuccess = function(data) {
  var response = JSON.parse(data.responseText);
  if (response.success == 1) {
    try {
      // hide the delete avatar trigger
      document.getElementById("trigger-delete-avatar").classList.add("hide");
      // unhide the camera icon in background
      document.querySelector(".dropzone-container-background").classList.remove("hide");
      // remove the profile picture from dropzone container
      document.querySelector(".dropzone-cotainer-profilepic").innerHTML = "";
      // remove the profile thumbnail from top nav bar and add icon
      document.querySelector(".top-nav-user-menu-trigger").innerHTML = "<i class='material-icons'>account_circle</i>";
      
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
//---------------------ENd DELETE AVATAR-----------------------//
//---------------------------------------------------------------//