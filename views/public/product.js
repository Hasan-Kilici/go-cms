let deletedProductTokens = "";

function addProductDeleteList(token){
    deletedProductTokens += token+"-";
}

function deleteSelectedProduct(){
    let ProductTokens = deletedProductTokens.slice(0, -1);
    let deletedProductsTokens = ProductTokens.split("-");
    let ProductCount = deletedProductsTokens.length;
    for(let i = 0;i < ProductCount;i++){
        deleteProduct(deletedProductsTokens[i])
    }
    let redirect = getCookie("LastPath")
}