var _instanceSweetalert = new sweetalert();
// init formData
var formData = {};
// init MODALS
var modalInstance;


document.addEventListener('DOMContentLoaded', function() {
  modalInstance = M.Modal.init(document.getElementById('modal-reusabletoatalqrcode'), {});
  // var elems = document.querySelectorAll(".modal");
  // var instances = M.Modal.init(elems, {});
});

function triggerReusableGenerateQrCodes(resuableId) {
  // set the reusable id to formData
	console.log(resuableId);

  formData["reusableid"] = resuableId;
  modalInstance.open();
}

function initAjaxReusableQrCode() {
  log.console("reusableNumberOfQrCodesElement")
  var reusableNumberOfQrCodesElement = document.getElementById("number_of_qr_codes").value;
  log.console(reusableNumberOfQrCodesElement)
  console.log(reusableNumberOfQrCodesElement)
  modalInstance.close();
  _instanceSweetalert.showBusy("Bitte warten ...");

  formData["number_of_qr_codes"] = reusableNumberOfQrCodesElement;
  setTimeout(function() {
	
    ajax("/reusables/addreusablesbatch", "POST", formData, handleCalculatedReusableSuccess);
  }, 500);
}

function handleCalculatedReusableSuccess(responseData) {
	
  _instanceSweetalert.close();
  const response = JSON.parse(responseData.response);
  if (response.error == 0) {
    setTimeout(function() {
      _instanceSweetalert.init("Aktion fehlgeschlagen!", response.errorMessage, true, "Ok", true);
      _instanceSweetalert.setType("error");
      _instanceSweetalert.show();
    }, 100);
  } else {
    
    setTimeout(function() {
      _instanceSweetalert.init("QR Codes!", "QR Codes wurden erfolgreich erzeugt.");
      _instanceSweetalert.setType("success");
      _instanceSweetalert.show();
      setTimeout(function() {
        _instanceSweetalert.close(1000);
      }, 100);
      // updateInvestorSection_1_4();
      scenarioSectionsForUpdateElement.innerHTML = response.parsedTemplate;
    }, 500);
  }
}


