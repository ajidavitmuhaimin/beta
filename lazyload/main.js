//example use IntersectionObserver on JavaScript

//define variable of IntersectionObserver and define callback on it

//get per one of element with for loop and insert it on entry
//get item when it on intersect area with if else
//and insert element what you want to detect with observer.observe()


var elementId=document.getElementById("element");
var observer=new IntersectionObserver((entries)=>{
 for(var i=0;i<entries.length;i++){
  var entry=entries[i];
  if(entry.isIntersecting){
    return true;
   }else{
    return false;
   }
 }
});
observer.observe(elementId);