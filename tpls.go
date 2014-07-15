package main

const templateList = `
<!DOCTYPE html>
<html lang="en" ng-app="fMgr">
  <head>
    <meta charset="utf-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">

    
    <!-- <link rel="stylesheet" href="//cdn.jsdelivr.net/codemirror/4.3.0/codemirror.css">
    <link rel="stylesheet" href="//cdn.jsdelivr.net/codemirror/4.3.0/theme/monokai.css">

    <link href="//netdna.bootstrapcdn.com/bootstrap/3.1.1/css/bootstrap.min.css" 
        rel="stylesheet" />
    <link rel="stylesheet" href="//cdnjs.cloudflare.com/ajax/libs/select2/3.5.0/select2.min.css" >
    
    <script src="//cdn.jsdelivr.net/codemirror/4.3.0/codemirror.js"></script>
    <script src="//cdn.jsdelivr.net/codemirror/4.3.0/addon/selection/active-line.js"></script>
    <script src="//cdn.jsdelivr.net/codemirror/4.3.0/keymap/sublime.js"></script>
    <script src="//cdn.jsdelivr.net/codemirror/4.3.0/addon/display/rulers.js"></script>
    <script src="//cdn.jsdelivr.net/codemirror/4.3.0/mode/css/css.js"></script>
    <script src="//cdn.jsdelivr.net/codemirror/4.3.0/mode/javascript/javascript.js"></script>
    <script src="//cdn.jsdelivr.net/codemirror/4.3.0/mode/markdown/markdown.js"></script>

    <script src="//code.jquery.com/jquery-1.11.0.min.js"> </script>
    <script src="//cdnjs.cloudflare.com/ajax/libs/select2/3.5.0/select2.min.js"> </script>
    <script src="//netdna.bootstrapcdn.com/bootstrap/3.1.1/js/bootstrap.min.js"> </script>
    <script src="//ajax.googleapis.com/ajax/libs/angularjs/1.2.19/angular.min.js"></script>
    <script src="//ajax.googleapis.com/ajax/libs/angularjs/1.2.19/angular-sanitize.js"></script>
    <script src="//cdn.jsdelivr.net/angular.bootstrap/0.11.0/ui-bootstrap-tpls.js"> </script>
    <script src="//cdn.jsdelivr.net/angular.bootstrap/0.11.0/ui-bootstrap-tpls.min.js"> </script>
    -->
    <link rel="stylesheet" href="/-/data/styles.css">
    <link rel="stylesheet" href="/-/data/app.css">
    <script src="/-/data/libs.js"></script>
    <script src="/-/data/app.js"></script>


    <style>
    

    </style>
<script>
var APP_PATH = "[% .Path %]"

</script>

  <title>FileManager</title>
  </head>
  <body ng-controller="ListCtr">
  

  <div class="container"  style="position:relative">
    <div class="row hidden-xs">
        <div class="col-md-6">
        <h3 id="title">FileManager</h3>
        </div>
    </div>
    
    <div  ng-show="flash.message" id="fmessage" class="{{ flash.type }}" >
    <p>{{ flash.message }}</p>
  </div>

  </div>
<!-- controller -->
<div class="container" id="nav">


<div class="row" id="header">

<div class="col-md-8" id="breadcrumb">
    <ol class="breadcrumb">
        <li><a href="/"><span class="glyphicon glyphicon-home"> </span></a></li>
        <li ng-repeat="item in Rutas"><a href="{{ item.url }}">{{ item.name }}</a></li> 
    </ol>
</div>

<div class="col-md-4" ng-controller="FinderCtrl" id="finder">
    <form role="form">
        <ui-select type="text" ui-select2="s2opts" 
            width="100%" ng-model="item.selected"
            >
        <ui-select-match placeholder="Directory quick finder">/{{$select.selected.path}}</ui-select-match>
        <ui-select-choices repeat="item in dirs"
             refresh="getData($select.search)"
             refresh-delay="0">
             <span class="glyphicon glyphicon-folder-close dir"></span> &nbsp; /<span ng-bind-html="item.path | highlight: $select.search"> </span>
        </ui-select-choices>
      <!--<div ng-bind-html="$select.search"></div>-->
    </ui-select-choices>
    </ui-select>
    </form>
</div>

</div>

</div>

<div class="container" id="list">

<div class="row" ng-show="view=='main'">
<form id="upload_files" style="display:none" enctype="multipart/form-data">
<input type="file" id="file_upload" name="files" multiple style="display:none" >
</form>
 <form role="form">
    <div class="col-md-6 col-xs-6">
        
        <inline-modal action="createFolder" btntext="Add Folder" 
            icon="glyphicon-folder-open" title="Folder Name:" handler="get_data()" 
            message="Folder created Created!" path="Path"></inline-modal>
        
        <button type="button" class="btn btn-info btn-sm" type="file" ng-click="AddFiles()"
        tooltip-placement="top" tooltip="Upload Multiple Files"><span class="glyphicon glyphicon-plus"> </span> <span class="hidden-xs">&nbsp; Upload</span></button>

        
        <inline-modal action="new_file" btntext="Add File" 
            icon="glyphicon-file" title="Filename:" handler="get_data()" 
            message="File Created!" path="Path"></inline-modal>

        &nbsp;&nbsp;&nbsp;
        
        <button class="btn btn-danger btn-sm" ng-show="selected>0" ng-click="DeleteSelected()"
            tooltip-placement="top" tooltip="Delete selected"><span class="glyphicon glyphicon-trash"> </span></button>

        

    </div>
    <div class="col-md-6 pull-right text-right  col-xs-6">
    
            
            <div class="btn-group-sm" style="float:right">
                <button type="button" class="btn ng-class: check_tipo('all');"  
                    tooltip-placement="top" tooltip="Show all"
                    ng-click="filter('all')">All</button>
                <button type="button" 
                    tooltip-placement="top" tooltip="Show only folders"
                    class="btn ng-class: check_tipo('folder');" ng-click="filter('folder')"><span class="glyphicon glyphicon-folder-open"> </span></button>
                <button type="button" 
                    tooltip-placement="top" tooltip="Show only files"
                    class="btn ng-class: check_tipo('files');" ng-click="filter('files')"><span class="glyphicon glyphicon-file"> </span></button>
                <button type="button" class="btn ng-class: check_tipo('hidden');"  ng-click="filter('hidden')" 
                    tooltip-placement="top" tooltip="Show Hidden"><span class="glyphicon glyphicon-eye-open"> </span></button>
            </div>
            
            <input ng-model="query" class="form-control input-sm" placeholder="Filter" style="width:45%; float:right; margin-right:15px" id="filter">

    </div>
    </form>
</div>


<div class="row list" ng-show="view=='main'" id="list">
<div class="col-md-12">
    <table class="table" ts-wrapper>
        <thead>
            <tr>
                <td  class="hidden-xs">&nbsp</td>
                <td ts-criteria="IsDir" class="f">+</td>
                <td ts-criteria="Name" ts-default="Ascending">Filename</td>
                <td ts-criteria="Size|parseFloat" class="hidden-xs">Size</td>
                <td ts-criteria="ModTime"  class="hidden-xs">Last Modified</td>
                <td><span class="hidden-xs">Actions</span></td>
            </tr>
        </thead>
        <tbody>
            <tr ng-repeat="item in Files|filter:query|filter:{'IsDir':ff}|filter:{'IsHidden':hidden}" ts-repeat>
                <td width="20" class="hidden-xs"><input type="checkbox" name="checkboxs[]" value="{{ item.Name }}" 
                    ng-click="CheckboxToggle(this, $event)" /></td>
                <td width="20"><span class="glyphicon glyphicon-folder-open" ng-show="item.IsDir"></span>
                        <span ng-hide="item.IsDir" class="glyphicon glyphicon-file"></span></td>
                
                <td inline-edit item="item" path="Path"> </td>
                
                <td width="100" class="hidden-xs">{{ item.Size|bytes }}</td>
                <td width="140" class="hidden-xs">{{ item.ModTime|date:'dd/MM/yyyy HH:mm:ss' }}</td>
                <td class="actions">
                
                    
                    

                    <!--<a href="{{ item.Name }}/?format=zip" target="_self" ng-if="item.IsDir" 
                        class="glyphicon glyphicon-download-alt delete" tooltip-placement="top" tooltip="Donwload as Zip"> </a>-->
                    
                    <div class="btn-group" dropdown>
                        <a href class="dropdown-toggle">
                            <span class="glyphicon glyphicon-align-justify"> </span>
                        </a>
                        <ul class="dropdown-menu" role="menu">
                            <li><a ng-click="CopyFile(item.Name)" href="#">Copy</a></li>
                            <li><a ng-click="DeleteFile(item.Name)" href="#">Delete</a></li>
                            <li ng-if="!item.IsDir"><a ng-click="EditFile(item.Name)" href="#">Edit</a></li>
                            <li><a ng-click="RenameFile(item.Name)" href="#">Rename</a></li>
                            <li><a ng-click="Compress(item.Name)" href="#">Compress</a></li>
                            <li ng-if="item.IsDir"><a href="{{ item.Name }}/?format=zip">Download as Zip</a></li>
                            
                        </ul>

                    </div>

                    <span ng-if="item.IsText" ng-click="EditFile(item.Name)" class="glyphicon glyphicon-pencil delete hidden-xs" tooltip-placement="top" tooltip="Edit" > </span>

                </td>
            </tr>
        </tbody>

    </table>
</div>
</div>

</div>


<div ng-show="view=='edit'" class="container">

<div class="row">
    <form role="form">
    <div class="col-md-6 col-xs-6">
        <p>{{ currentEditedFile }}</p>
    </div>
    <div class="col-md-6 pull-right text-right col-xs-6">
        <button type="button" class="btn btn-info btn-sm" ng-click="ToView('main')"><span class="glyphicon glyphicon-arrow-left"> </span><sapn class="hidden-xs"> &nbsp; Back</sapn></button>
        <button type="button" class="btn btn-info btn-sm" ng-click="SaveFile()"><span class="glyphicon glyphicon-floppy-disk"> </span><span class="hidden-xs"> &nbsp; Save</span></button>
        
    </div>
    </form>
</div>

<div class="row" class="mtop" style="margin-top:20px">
    <div class="col-md-12">
        <textarea id="editor" ui-codemirror="{onLoad:codemirrorLoaded}"
            ui-codemirror-opts="editorOptions"
            ng-model="EditorCurrentContent"
            ui-refresh='EditorRefresh'>
        </textarea>
    </div>
</div>

</div>


<div class="container">
    <div class="row" style="margin-top:50px; border-top:1px solid #eaeaea; padding-top:20px; font-size:10px">
        <div class="col-md-12">
        <p><a href="http://github.com/jordic/file_server">http://github.com/jordic/file_server</a> -- v.[% .version %]
        </p></div>
    </div>
</div>
    

  </body>

<script type="text/ng-template" id="/inline-modal.html">
    <button class="btn btn-info btn-sm" ng-click="ClickButton()">
        <span class="glyphicon {{ icon }}"> </span><span class="hidden-xs">&nbsp; {{ btntext }}</span></button>

    <div ng-show="folder_popover==true" class="popover bottom am-flip-x" 
        style="display: block; width:200px">
                <div class="arrow"></div>
                <h3 class="popover-title">{{ title }}</h3>
                <div class="popover-content">
                    <input type="text" class="form-control input-sm" 
                        ng-model="filename" id="folder_field" 
                        tfocus="folder_popover">
                    <button type="button" ng-click="Process()" 
                        class="btn btn-info btn-sm pull-right" 
                            style="margin-top:5px; margin-bottom:5px">Add</button>
                </div>
        </div>


</script>

<script type="text/ng-template" id="/inline-edit.html">
    <td ng-mouseenter="roll=true" ng-mouseleave="roll=false" class="tdrename"
        ng-click="Show($event)">    
        <a href="{{ Path }}{{ item.Name }}" target="_self" 
            ng-if="!item.IsDir">{{ item.Name }}</a>
        <a  ng-click="Go(item.Name)" ng-if="item.IsDir" class="dir">{{ item.Name }}</a><span class="visible-xs small">{{ item.ModTime|date:'dd/MM/yyyy HH:mm:ss' }} | {{ item.Size|bytes }}</span>
    

    <div class="col-md-12 renameinput inline-form" ng-show="showrename==true">
        <input type="text" class="" name="value" 
            ng-model="item.Name" tfocus="showrename" />
        <button class="btn btn-info btn-xs" ng-click="SaveItem()">SAVE</button>
        <button class="btn btn-default btn-xs" ng-click="showrename=false">CANCEL</button>
    </div>
    </td>
</script>

<script>









</script>




</html>
`
