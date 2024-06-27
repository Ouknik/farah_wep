
// init formData
var formData = {};
// init MODALS
var modalInstance;
var _instanceSweetalert = new sweetalert();
document.addEventListener('DOMContentLoaded', function() {
    modalInstance = M.Modal.init(document.getElementById('modal-reusabletoatalqrcode'), {});
  // var elems = document.querySelectorAll(".modal");
  // var instances = M.Modal.init(elems, {});
});


function triggerReusableGenerateQrCodes(resuableId) {
	// set the reusable id to formData
	formData["reusableid"] =resuableId
	modalInstance.open();
}






function initAjaxReusableQrCode() {
    var reusableNumberOfQrCodesElement =document.getElementById("number_of_qr_codes").value;
	modalInstance.close();
    _instanceSweetalert.showBusy("Bitte warten ...");
	
	formData["number_of_qr_codes"] = reusableNumberOfQrCodesElement
	
    ajax("/reusables/addreusablesbatch", "post",formData , handleCalculated_reusable_success);  
	
}



function handleCalculated_reusable_success(responseData) {
	
    const response = JSON.parse(responseData.response);
	_instanceSweetalert.close();
	if (response.error == 1) {
        setTimeout(function(){
            _instanceSweetalert.init( response.message, true, "Ok", true);
            _instanceSweetalert.setType("error");
            _instanceSweetalert.show();
        }, 100);
    } else {
		
		setTimeout(
            function() {
                _instanceSweetalert.init(response.message, "");
                _instanceSweetalert.setType("success");
                _instanceSweetalert.show();
                setTimeout(function () {
                    _instanceSweetalert.close(1000);
                }, 100);
                // updateInvestorSection_1_4();
                scenarioSectionsForUpdateElement.innerHTML = response.parsedTemplate

            },
        500);

	}
}

