$(function () {
  
    loadMainMenu();
    
   
})


function loadMainMenu() {
    $('.easyui-tree').tree({
        onClick: function (node) {
            if ($('.easyui-tree').tree('isLeaf', node.target)) { 
                //判断要打开的页面是否已经打开
                //若打开 则选中该页面 没有就新建一个页面           
                if ($('#tabs').tabs('exists', node.text)) {
                    //若打开 则选中该页面 
                    $('#tabs').tabs('select',  node.text);                 
                }else{
                    //就新建一个页面    
                    addTab(node.text,node.attributes.url);
                }
                
            } else {
                //
                $(this).tree(node.state === 'closed' ? 'expand' : 'collapse', node.target);
            }
    
    
            // alert node text property when clicked
        }
    });
}

function addTab(title, url) {   
    var iframe ='<iframe src="'+url+'" frameborder="0" style="height: 100%;width: 100%;" scrolling="yes"></iframe>';
        $('#tabs').tabs('add', {
                title: title,
                closable: true,
                href:url
                //content:iframe
            });
   
}



