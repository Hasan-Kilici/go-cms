let deletedItemsTokens = "";

function addBlogDeleteList(token){
    deletedItemsTokens += token+"-";
}

function deleteSelectedBlogs(){
    let blogTokens = deletedItemsTokens.slice(0, -1);
    let deletedBlogTokens = blogTokens.split("-");
    let blogCount = deletedBlogTokens.length;
    for(let i = 0;i < blogCount;i++){
        deleteBlog(deletedBlogTokens[i])
    }
    let redirect = getCookie("LastPath")
}