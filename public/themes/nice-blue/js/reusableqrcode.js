// init formData
var formData = {};

var _instanceSweetalert = new sweetalert();



function DeleteReusableQrCode( ReusableId) {
  console.log(select)
  console.log(ReusableId)
   
    formData["reusableId"]  = ReusableId;
    _instanceSweetalert.showBusy("Bitte warten ...");
   // ajax("/reusables/addBatch", "post",formData , handleCalculated_reusable_success);  
   setTimeout(
    function() {
        ajax("/reusables/qrcode/changestatue", "POST", formData, handleSuccessChangeStatus);
    },500
);
       
  }


  const scenarioSectionsForUpdateElement = document.getElementById("switch-blocked");
  function handleSuccessChangeStatus(responseData) {
      const response = JSON.parse(responseData.response);
      _instanceSweetalert.close();
       console.log(response)
      if (response.error == 1) {
          setTimeout(function(){
              _instanceSweetalert.init("Der Vorgang wurde erfolgreich abgeschlossen!", response.errorMessage, true, "Ok", true);
              _instanceSweetalert.setType("error");
              _instanceSweetalert.show();
          }, 100);
      } else {
  
          setTimeout(
              function() {
                  _instanceSweetalert.init("Erfolgreich blockiert!", "");
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

