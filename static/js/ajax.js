/**
 * ajax 函数
 * 参数为对象 包含量 方法 type  get post 
 *  url 
 * data (key1=value1&key2=value2)
 * success 回调函数
 */
function ajax(option){
    var xhr=new XMLHttpRequest();
    if(option.type=='get'&&option.data){
        option.url+='?';
        option.url+=data;
        option.data=null;
    }
    xhr.open(option.type,option.url);
    if(option.type=='post'&&option.data){
        xhr.setRequestHeader('content-type','application/x-www-form-urlencoded');
    }    
    xhr.onreadystatechange=function(){        
        if(xhr.readyState==4&&xhr.status==200){
            var type=xhr.getResponseHeader('content-type');
            //判断返回的数据是不是json
            if(type.indexOf('json')!=-1){
                option.success(JSON.parse(xhr.responseText));
                //判断返回的数据是不是xml
            }else if(type.indexOf('xml')!=-1){
                option.success(xhr.responseXML);
            }else{//返回的数据不是json 也不是xml 普通字符串
                option.success(xhr.responseText);
            }
        }
       
    }
    xhr.send(option.data);
}