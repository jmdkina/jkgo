/**
 * Created by jmdvi on 2017/7/22.
 */

$(function(){
    Vue.config.delimiters = ['${', '}#'];
   var add_fileserver = new Vue({
       el: '#addfileserver',
       data: {
           path: "",
           uploadpath: "",
           result:""
       },
       methods: {
           addfileserver: function() {
               var str = "jk=addfileserver&path=" + this.path;
               console.log(str);
               this.result = "";
               $.ajax({
                   url: "dir",
                   method:"POST",
                   data: str,
                   success: function() {
                       console.log("addfileserver success");
                       add_fileserver.result = "success";
                   }
               });
           },
           adduploadpath: function() {
               var str = "jk=adduploadpath&path=" + this.uploadpath;
               console.log("adduploadpath " + str);
               $.ajax({
                   url: "dir",
                   method:"POST",
                   data: str,
                   success: function() {
                       console.log("add uploadpath success");
                   }
               })
           }
       }
    });

});