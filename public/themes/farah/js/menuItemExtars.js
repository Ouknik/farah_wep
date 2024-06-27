//global variable 
var divSerachelmFree = document.getElementById("resultat-search-free");     
var divSerachelmPaid = document.getElementById("resultat-search-paid");     
var inputSearchElmFree = document.getElementById("input-search-free");     
var inputSearchElmPaid = document.getElementById("input-search-paid");    
var spinnerElm =  document.getElementById("spinner") 
var errorDiv = document.getElementById("errorr-message")
document.addEventListener("DOMContentLoaded", function () {
  var elems = document.querySelectorAll(".modal");
  var instances = M.Modal.init(elems, {});
  console.log(instances)
});

function searchToExtrasItem(e,type) {
  typeOfExtrasMenus = type
  
  if (e.value !="" && e.value !=null){
    show(spinnerElm)
    ajax("/menuitemextras/search","POST", { type_search: typeOfExtrasMenus, value_search: e.value,item_id:menuItemId },SearchExtrasmenuSuccess);
 
  } else{
    divSerachelmFree.innerHTML = "";
    divSerachelmPaid.innerHTML =""
  }

  
}

SearchExtrasmenuSuccess = function (data) {
  hide(spinnerElm)
  var jsonData = JSON.parse(data.response);
  divSerachelmFree.innerHTML = "";
  divSerachelmPaid.innerHTML = "";
  
  Object.keys(jsonData).forEach((key) => {addRow(jsonData[key]); });
};


function addRow(etrastItem) {
    
  const div = document.createElement("div");
  var newprice = etrastItem.extra_item_price == "0"?"":etrastItem.extra_item_price;
  var divSerachelm =etrastItem.extra_item_price == "0"?divSerachelmFree : divSerachelmPaid;
  div.className = "row";

  div.innerHTML =
    ` <p>
    <label onclick="event.preventDefault(); addMenuItemExtrat(`+etrastItem.extra_item_id+`,'`+etrastItem.extra_item_name+`')">
      <input type="radio" value="`+etrastItem.extra_item_id+`" />
      <span>`+etrastItem.extra_item_name+` |</span>
      <span>`+etrastItem.extra_item_name_display+`</span>
      <span>`+newprice +`</span>
    </label>
  </p>
  `;

  divSerachelm.appendChild(div);
}

function addMenuItemExtrat(id,name ){


  if(menuItemId == undefined || menuItemId == null){
    errorMessage("You_just_can_to_add_a_cost_or_free_items_on_edit_mode")
  
  }else{
    show(spinnerElm)
    ajax("/menuitemextras/add","POST", { item_id: menuItemId, extra_item_id: id, extra_type:typeOfExtrasMenus,extras_item_name:name },addExtrasmenuSuccess);
  }
  inputSearchElmFree.value =""
  inputSearchElmPaid.value =""
  divSerachelmFree.innerHTML=""
  divSerachelmPaid.innerHTML = ""

  
  

}


// window.onload  = function (){
//   if(menuItemId == undefined || menuItemId == null){
//     //alert("You just can to add a cost or free items on edit mode !")
//   }else{
    
//   }
// }
function loadExtrasItems(id){
  console.log(id,menuItemId)
  if(menuItemId == undefined || menuItemId == null){
    //     //alert("You just can to add a cost or free items on edit mode !")
  }else{
      loadTheItemMenuFromServer()
  }
  
}
function loadTheItemMenuFromServer(){
  document.getElementById("free-extras" ).innerHTML=""
  document.getElementById("paid-extras").innerHTML=""
  show(spinnerElm)
  ajax("/menuitemextras/","POST", { item_id: menuItemId },LoadMenuItemExtratSucces);
}
function LoadMenuItemExtratSucces(data){
  hide(spinnerElm)
  var responseData =  JSON.parse(data.response)
  Object.keys(responseData).forEach((key) => {
    addListToPaidOrFreeExtras(responseData[key]);
  });
}
function addExtrasmenuSuccess(data){
  
  hide(spinnerElm)
  var responseData =  JSON.parse(data.response)
  if(responseData.errorMsg){
    errorMessage(responseData.errorMsg)
  }else{
    document.getElementById("free-extras" ).innerHTML=""
    document.getElementById("paid-extras").innerHTML=""
   
    Object.keys(responseData).forEach((key) => {
      addListToPaidOrFreeExtras(responseData[key]);
    });
  }




}

function errorMessage(message){
  errorDiv.style.display ="block"
  errorDiv.innerHTML = message
  // 5s to show the error message
  setTimeout(()=>{
    errorDiv.innerHTML = ""
    errorDiv.style.display = "none";
  },5000)

}
function addListToPaidOrFreeExtras(extras){
  //choise ul item for add the extras if 0 is free is'nt paid
  var idElement  = extras.extra_type == "0" ? "free-extras" :"paid-extras"
  var elementFreeOrpaid =  document.getElementById(idElement)
  const li = document.createElement("li");
  li.classList.add("collection-item")
  li.innerHTML =
    ``+extras.extra_item_name+`
    <a style="float:right"  onclick="deleteItemExtars(`+extras.extra_id+`)" class="  waves-effect waves-light " >
      <i class="material-icons">delete</i>
    </a>
      
  `;
  elementFreeOrpaid.appendChild(li)
}


function deleteItemExtars(id){
  //dis method just for delete we not need now because not know if delete or just update 
// if 
  if(id){
    ajax("/menuitemextras/delete","POST", { id: id},loadTheItemMenuFromServer);

  }
}


function hide(elem){
  elem.style.display = 'none';
}

function show(elem){
  elem.style.display = 'block';
}




/*//global variable 
var divSerachelm = document.getElementById("resultat-search");     
var errorDiv = document.getElementById("errorr-message")

document.addEventListener("DOMContentLoaded", function () {
  var elems = document.querySelectorAll(".modal");
  var instances = M.Modal.init(elems, {});
});

function addToListExtras(e) {
  if (e.value !="" && e.value !=null){
    ajax("/menuitemextras/search","POST", { type_search: typeOfExtrasMenus, value_search: e.value },SearchExtrasmenuSuccess);
  } else{
    divSerachelm.innerHTML = "";
  }

  
}

SearchExtrasmenuSuccess = function (data) {
  var jsonData = JSON.parse(data.response);
  divSerachelm.innerHTML = "";
  Object.keys(jsonData).forEach((key) => {addRow(jsonData[key]); });
};

function addRow(etrastItem) {
    
  const div = document.createElement("div");
  var newprice = etrastItem.extra_item_price == "0"?"":etrastItem.extra_item_price;
  div.className = "row";

  div.innerHTML =
    ` <p>
    <label onclick="event.preventDefault(); addMenuItemExtrat(`+etrastItem.extra_item_id+`,'`+etrastItem.extra_item_name+`')">
      <input type="radio" value="`+etrastItem.extra_item_id+`" />
      <span>`+etrastItem.extra_item_name+` |</span>
      <span>`+etrastItem.extra_item_name_display+`</span>
      <span>`+newprice +`</span>
    </label>
  </p>
  `;

  divSerachelm.appendChild(div);
}

function addMenuItemExtrat(id,name ){


  if(menuItemId == undefined || menuItemId == null){
    alert("You just can to add a cost or free items on edit mode !")
  }else{
    ajax("/menuitemextras/add","POST", { item_id: menuItemId, extra_item_id: id, extra_type:typeOfExtrasMenus,extras_item_name:name },addExtrasmenuSuccess);
  }
  

}


window.onload  = function (){
  if(menuItemId == undefined || menuItemId == null){
    alert("You just can to add a cost or free items on edit mode !")
  }else{
    loadTheItemMenuFromServer()
  }
}

function loadTheItemMenuFromServer(){
  document.getElementById("free-extras" ).innerHTML=""
  document.getElementById("paid-extras").innerHTML=""
  ajax("/menuitemextras/","POST", { item_id: menuItemId },LoadMenuItemExtratSucces);
}

function LoadMenuItemExtratSucces(data){
  var responseData =  JSON.parse(data.response)
  Object.keys(responseData).forEach((key) => {
    addListToPaidOrFreeExtras(responseData[key]);
  });
}

function addExtrasmenuSuccess(data){
  

  var responseData =  JSON.parse(data.response)
  if(responseData.errorMsg){
    errorDiv.style.display ="block"
    errorDiv.innerHTML = responseData.errorMsg
  }else{
    document.getElementById("free-extras" ).innerHTML=""
    document.getElementById("paid-extras").innerHTML=""
   
    Object.keys(responseData).forEach((key) => {
      addListToPaidOrFreeExtras(responseData[key]);
    });
  }



// 5s to show the error message
  setTimeout(()=>{
    errorDiv.innerHTML = ""
    errorDiv.style.display = "none";
  },5000)
}

function addListToPaidOrFreeExtras (extras){
  //choise ul item for add the extras if 0 is free is'nt paid
  var idElement  = extras.extra_type == "0" ? "free-extras" :"paid-extras"
  var elementFreeOrpaid =  document.getElementById(idElement)
  const li = document.createElement("li");
  li.classList.add("collection-item")
  li.innerHTML =
    ``+extras.extra_item_name+`
    <a style="float:right"  onclick="deleteItemExtars(`+extras.extra_id+`)" class="  waves-effect waves-light " >
      <i class="material-icons">delete</i>
    </a>
      
  `;
  elementFreeOrpaid.appendChild(li)
}


function deleteItemExtars(id){
  //dis method just for delete we not need now because not know if delete or just update 
// if 
  if(id){
    ajax("/menuitemextras/delete","POST", { id: id},loadTheItemMenuFromServer);

  }
}




*/