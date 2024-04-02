/**Variables */

var languagedata;

let editor;

let ckeditor1;

let collapsecount = 1

let newPage = 0

let OrderIndex = 1

// let suborderindex = 1

let PageId = 1

let SubPageId = 1

let GroupId = 1

let newpagesarr = []

let newGrouparr = []

let Subpagearr = []

let NewSubObj = []

let overallarray = []

/**delete array*/
let deletePagearr = []

let deletegrouparr = []

let deletesubarr = []

/** */
$(document).ready(async function () {

    var languagepath = $('.language-group>button').attr('data-path')
  
    await $.getJSON(languagepath, function (data) {
        
        languagedata = data
    })

})


$(document).ready(function () {

    CKEDITORS()

    $.ajax({
        type: "get",
        url: "/spaces/page/pagedata",
        dataType: 'json',
        data: {
            sid: $("#spid").val(),
        },
        cache: false,
        success: function (result) {
         
        }
    })

    if (newpagesarr.length >= 0) {

        var pagename = "Introduction"

        var pagedata = AddPageString(pagename, PageId)

        $('.newpages').append(pagedata)

        var pageobj = PageObjectCreate();

        pageobj.Name = pagename

        newpagesarr.push(pageobj)

        PageId++;

        OrderIndex++;

    }

    $('.viewpage').each(function(){

        $(this).children('.section-fields-content-child').trigger('click');

        return;
    })

   
})

function CKEDITORS() {

    CKEDITOR.ClassicEditor.create(document.getElementById("neweditor"), {
        toolbar: {
            items: ['heading', 'bold', 'italic', 'alignment', 'underline', 'blockQuote', 'undo', 'redo', 'imageStyle:wrapText', 'imageStyle:breakText', 'imageUpload', 'indent', 'outdent', 'numberedList', 'bulletedList', 'todoList'],
            shouldNotGroupWhenFull: true
        },
        list: {
            properties: {
                styles: true,
                startIndex: true,
                reversed: true
            }
        },
        heading: {
            options: [
                { model: 'paragraph', title: 'Paragraph', class: 'ck-heading_paragraph' },
                { model: 'heading1', view: 'h1', title: 'Heading 1', class: 'ck-heading_heading1' },
                { model: 'heading2', view: 'h2', title: 'Heading 2', class: 'ck-heading_heading2' },
                { model: 'heading3', view: 'h3', title: 'Heading 3', class: 'ck-heading_heading3' },
                { model: 'heading4', view: 'h4', title: 'Heading 4', class: 'ck-heading_heading4' },
                { model: 'heading5', view: 'h5', title: 'Heading 5', class: 'ck-heading_heading5' },
                { model: 'heading6', view: 'h6', title: 'Heading 6', class: 'ck-heading_heading6' }
            ]
        },
        placeholder: 'This is some sample Content',
        fontFamily: {
            options: [
                'default',
                'Arial, Helvetica, sans-serif',
                'Courier New, Courier, monospace',
                'Georgia, serif',
                'Lucida Sans Unicode, Lucida Grande, sans-serif',
                'Tahoma, Geneva, sans-serif',
                'Times New Roman, Times, serif',
                'Trebuchet MS, Helvetica, sans-serif',
                'Verdana, Geneva, sans-serif'
            ],
            supportAllValues: true
        },
        fontSize: {
            options: [10, 12, 14, 'default', 18, 20, 22],
            supportAllValues: true
        },
        htmlSupport: {
            allow: [
                {
                    name: /.*/,
                    attributes: true,
                    classes: true,
                    styles: true
                }
            ]
        },
        htmlEmbed: {
            showPreviews: true
        },
        link: {
            decorators: {
                addTargetToExternalLinks: true,
                defaultProtocol: 'https://',
                toggleDownloadable: {
                    mode: 'manual',
                    label: 'Downloadable',
                    attributes: {
                        download: 'file'
                    }
                }
            }
        },
        mention: {
            feeds: [
                {
                    marker: '@',
                    feed: [
                        '@apple', '@bears', '@brownie', '@cake', '@cake', '@candy', '@canes', '@chocolate', '@cookie', '@cotton', '@cream',
                        '@cupcake', '@danish', '@donut', '@dragée', '@fruitcake', '@gingerbread', '@gummi', '@ice', '@jelly-o',
                        '@liquorice', '@macaroon', '@marzipan', '@oat', '@pie', '@plum', '@pudding', '@sesame', '@snaps', '@soufflé',
                        '@sugar', '@sweet', '@topping', '@wafer'
                    ],
                    minimumCharacters: 1
                }
            ]
        },
        removePlugins: [
            'ExportPdf',
            'ExportWord',
            'CKBox',
            'CKFinder',
            'EasyImage',
            'RealTimeCollaborativeComments',
            'RealTimeCollaborativeTrackChanges',
            'RealTimeCollaborativeRevisionHistory',
            'PresenceList',
            'Comments',
            'TrackChanges',
            'TrackChangesData',
            'RevisionHistory',
            'Pagination',
            'WProofreader',
            'MathType',
            'SlashCommand',
            'Template',
            'DocumentOutline',
            'FormatPainter',
            'TableOfContents',
            'PasteFromOfficeEnhanced'
        ]
    }).then(ckeditor => {
        ckeditor1 = ckeditor
        ckeditor.plugins.get('FileRepository').createUploadAdapter = (loader) => {
            return {
                upload: () => {
                    return loader.file.then(file => {
                        return new Promise((resolve, reject) => {
                            const formData = new FormData();
                            formData.append('file', file);
                            formData.append('csrf', $("input[name='csrf']").val())
                            fetch('/channel/imageupload', {
                                method: 'POST',
                                body: formData
                            })
                                .then(response => response.json())
                                .then(data => {
                                    if (data.success) {
                                        resolve({
                                            default: data.url // URL to the uploaded image
                                        });
                                    } else {
                                        reject(data.error);
                                    }
                                })
                                .catch(error => {
                                    reject(error);
                                });
                        });
                    });
                }
            };
        };
    })
        .catch(error => {
            console.error(error);
        });

}

/** Page String */
function AddPageString(name, newid) {

    var pagestr = `
<div class="sort-index viewpage page0${newid}" data-id=0 data-newid=${newid}>
    <a class="section-fields-content-child" href="javascript:void(0)">
        <h3 class="heading-three">${name}</h3>
        <div class="card-option field-delete">
            <div class="btn-group language-group space-option-grp">
                <button type="button" class="dropdown-tog" data-bs-toggle="dropdown" aria-expanded="true">
                    <img src="/public/img/card-option.svg" alt="">
                    <img src="/public/img/dark-option.svg" alt="">
                    <img src="/public/img/option-primary.svg" alt="">
                </button>

                <ul class="dropdown-menu dropdown-menu-end" data-popper-placement="bottom-end" style="position: absolute;inset: 0px 0px auto auto;margin: 0px;transform: translate(0px, 22px);">
                    <li>
                        <button class="dropdown-item insertsubpage" type="button">
                            Insert Subpage
                        </button>
                    </li>
                    <li>
                        <button class="dropdown-item pageabove" type="button">
                            Add Page Above
                        </button>
                    </li>

                    <li>
                        <button class="dropdown-item pagebelow" type="button">
                            Add Page Below
                        </button>
                    </li>
                    <li>
                        <button class="dropdown-item rename" type="button">
                            Rename
                        </button>
                    </li>
                    <li>
                        <button class="dropdown-item pageclone" type="button">
                            Clone
                        </button>
                    </li>
                    <li>
                        <button class="dropdown-item movepage" type="button">
                            Move Page
                        </button>
                    </li>
                    <li>
                        <button class="dropdown-item deletepage" type="button" data-bs-toggle="modal" data-bs-target="#centerModal">
                            Delete
                        </button>
                    </li>
                </ul>
            </div>
        </div>
    </a>
</div>`

    return pagestr
}

function AddGroupString() {

}

function AddSubString() {

}

/** View Page content */
$(document).on('click', '.viewpage', function () {

    var id = $(this).attr('data-id');

    var newid = $(this).attr('data-newid');

    $('.section-fields-content-child').removeClass('active').removeClass('pageselected');

    $(this).children('a').addClass('active').addClass('pageselected');


    const pageindex = newpagesarr.findIndex(obj => {

        return obj.PageId == id && obj.NewPageId == newid;

    })


    if (pageindex >= 0) {

        $('#title').val(newpagesarr[pageindex].Name)

        ckeditor1.setData(newpagesarr[pageindex].Content);

    }

})

/** Keyup content save to object */
$(document).on("DOMSubtreeModified", ".ck-restricted-editing_mode_standard", function () {
    var data = ckeditor1.getData();
    var pageid = $('.pageselected').parents('.viewpage').attr('data-id');
    var newpageid = $('.pageselected').parents('.viewpage').attr('data-newid');

    // if (pgeid == undefined) {
    //     var subid = $('.pageselected').parents('.accordion-collapse').attr('data-sub-id')
    //     var newpgid = $('.pageselected').parents('.accordion-collapse').attr('data-newid')
    //     var newparentid = $('.pageselected').parents('.accordion-collapse').attr('data-parent-id')
    //     for (let x of Subpage) {
    //         if (x['ParentId'] == newparentid && x['SpgId'] == subid && x['NewSpId']==newpgid) {
    //             x['Content'] = data
    //             break
    //         }
    //     }

    // } else {

        const pageindex = newpagesarr.findIndex(obj => {

            return obj.PageId == pageid && obj.NewPageId == newpageid;
    
        })
    
        if (pageindex >= 0) {
            
            newpagesarr[pageindex].Content = data

        }    
       
    // }
})

/**Page above  */
$(document).on('click', '.pageabove', function () {

    var pagename = "New Page " + PageId

    var AddPage = AddPageString(pagename, PageId);

    $(AddPage).insertBefore($(this).parents(".sort-index"));

    var pageobj = PageObjectCreate();

    pageobj.Name = pagename

    newpagesarr.push(pageobj)

    PageId++;

    OrderIndex++;
})

/**Page below  */
$(document).on('click', '.pagebelow', function () {

    var pagename = "New Page " + PageId

    var AddPage = AddPageString(pagename, PageId);

    $(AddPage).insertAfter($(this).parents(".sort-index"));

    var pageobj = PageObjectCreate();

    pageobj.Name = pagename

    newpagesarr.push(pageobj)

    PageId++;

    OrderIndex++;
})

/** Clone Page */
$(document).on('click', '.pageclone', function () {

    var id = $(this).parents('.sort-index').attr('data-id');

    var newid = $(this).parents('.sort-index').attr('data-newid');

    const pageindex = newpagesarr.findIndex(obj => {

        return obj.PageId == id && obj.NewPageId == newid;

    })

    if (pageindex >= 0) {

        var AddPage = AddPageString(newpagesarr[pageindex].Name, PageId);

        $(AddPage).insertAfter($(this).parents(".sort-index"));

        var pageobj = PageObjectCreate();

        pageobj.Name = newpagesarr[pageindex].Name

        newpagesarr.push(pageobj)

        PageId++;

        OrderIndex++;
    }

})

/** keyup name change */
$(document).on('keyup', ".pagetitle", function () {
    
    var id = $('.pageselected').parents(".viewpage").attr('data-id');
    
    var newid = $('.pageselected').parents(".viewpage").attr('data-newid');
    
    var value = $(this).val();
    
    $('.pageselected').children('h3').text(value);

    const pageindex = newpagesarr.findIndex(obj => {

        return obj.PageId == id && obj.NewPageId == newid;

    })

    if (pageindex >= 0) {

        newpagesarr[pageindex].Name= value

    }

})

/*Delete page Confirmation */
$(document).on('click', '.deletepage', function () {

    var pageid = $(this).parents(".sort-index").attr("data-id")

    var newpgid = $(this).parents(".sort-index").attr("data-newid")

    const pageindex = newpagesarr.findIndex(obj => {

        return obj.PageId == pageid && obj.NewPageId == newpgid;

    })

    if (pageindex >= 0) {

        $("#delid").attr("data-pageid", pageid)

        $('#delid').attr('data-newpageid', newpgid)

        $('.deltitle').text('Delete Page ?')

        $('.deldesc').text('Are you sure ! you want to delete this Page?')

        $('.delname').text(newpagesarr[pageindex].Name)

        $('#content').text("Are you sure! you want to delete this Page?")

    }

})

/** Delete Page & Group & Subpage */
$(document).on('click', '#delid', function () {

    var pageid = $(this).attr('data-pageid');

    var newpageid = $(this).attr('data-newpageid');

    var groupid = $(this).attr('data-groupid');

    var newgroupid = $(this).attr('data-newgroupid');

    var subpageid = $(this).attr('data-subpageid');

    var newsubpageid = $(this).attr('data-newsubpageid');

    if (pageid != undefined || newpageid != undefined) {

        for (let x of Subpagearr) {

            if (x.ParentId) {

            }

        }

  

        const pageindex = newpagesarr.findIndex(obj => {

            console.log(obj.PageId,pageid,obj.NewPageId,newpageid);

            return obj.PageId == pageid && obj.NewPageId == newpageid;

        })

        if (pageindex >= 0) {

            console.log("asfd");

            newpagesarr.splice(pageindex, 1)

            $('.page'+pageid+newpageid).remove();

        }

    }

    if (groupid != undefined || newgroupid != undefined) {

        const groupindex = newGrouparr.findIndex(obj => {

            return obj.GroupId == groupid && obj.NewGroupId == newgroupid;

        })

        if (groupindex >= 0) {

            newGrouparr.splice(groupindex, 1)

        }

    }

    $('#delcancel').trigger('click');
})



/**Add Page */
$(document).on('click', '#addpage', function () {

    var pagename = "New Page " + PageId

    var pagedata = AddPageString(pagename, PageId)

    $('.newpages').append(pagedata)

    var pageobj = PageObjectCreate();

    pageobj.Name = pagename

    newpagesarr.push(pageobj)

    PageId++;

    OrderIndex++;

})

/** Page Creation Object */
function PageObjectCreate() {

    var obj = {}

    obj.PageId = 0,

    obj.NewPageId = PageId,

    obj.Content = "",

    obj.PageGroupId = 0,

    // obj.NewGroupId=parseInt(newgrpid==undefined?0:newgrpid),

    obj.OrderIndex = OrderIndex,

    obj.ParentId = 0

    return obj

}


/** SubPage Creation Object */
function SubPageObjectCreate() {

    var obj = {}

    obj.SubPageId = 0,

    obj.NewSubPageId = SubPageId,

    obj.Content = "",

    obj.PageGroupId = 0,

    obj.NewGroupId = parseInt(newgrpid == undefined ? 0 : newgrpid),

    obj.OrderIndex = orderindex

    obj.ParentId= parseInt(parentid)

    return obj

}


/** Group Creation Object */
function GroupCreationObject() {

    var obj = {}

    obj.GroupId = 0,

    obj.NewGroupId = GroupId,

    // obj.Name = name.toUpperCase(),

    obj.OrderIndex = orderindex

    return obj

}


function PGList() {


    for (let x of overallarray) {

        OrderIndex = x['OrderIndex']

        /**this page */
        if (x['PgId'] != 0 && x['NewPgId'] == 0 && x['Pgroupid'] == 0 && x['NewGrpId'] == 0) {

            var AddPage = AddPageString(x['Name'], x['PgId'], 0);

            $('.newpages').append(AddPage);

            for (let j of Subpage) {

                if (j['ParentId'] == x['PgId'] && j['NewParentId'] == 0) {

                    var AddSubPage = AddSubPageString(j['Name'], x['ParentId'] == 0 ? x['NewParentId'] : x['ParentId'], j['SpgId'], "", j['NewParentId'])

                    $('#accordionExample' + j['ParentId']).children('.accordion-item').append(AddSubPage)

                } else if (j['ParentId'] == x['NewPgId'] && j['NewParentId'] == 0) {

                    var AddSubPage = AddSubPageString(j['Name'], x['ParentId'] == 0 ? x['NewParentId'] : x['ParentId'], j['SpgId'], "", j['NewParentId'])

                    $('#accordionExample' + j['ParentId']).children('.accordion-item').append(AddSubPage)

                } else if (j['NewParentId'] == x['NewPgId'] && j['ParentId'] == 0) {

                    var AddSubPage = AddSubPageString(j['Name'], x['ParentId'] == 0 ? x['NewParentId'] : x['ParentId'], j['SpgId'], "", 0)

                    $('#accordionExample' + j['NewParentId']).children('.accordion-item').append(AddSubPage)

                } else if (j['NewParentId'] == x['PgId'] && j['ParentId'] == 0) {

                    var AddSubPage = AddSubPageString(j['Name'], x['ParentId'] == 0 ? x['NewParentId'] : x['ParentId'], j['SpgId'], "", j['NewParentId'])

                    $('#accordionExample' + j['NewParentId']).children('.accordion-item').append(AddSubPage)

                }

            }


        } else if (x['NewPgId'] !== 0 && x['PgId'] == 0 && x['Pgroupid'] == 0 && x['NewGrpId'] == 0) {

            var AddPage = AddPageString(x['Name'], 0, x['NewPgId']);

            $('.newpages').append(AddPage);

            for (let j of Subpage) {

                if (j['ParentId'] == x['PgId']) {

                    var AddSubPage = AddSubPageString(j['Name'], x['ParentId'], j['SpgId'], "", 0)

                    $('#accordionExample' + j['ParentId']).children('.accordion-item').append(AddSubPage)

                } else if (j['ParentId'] == x['NewPgId']) {

                    var AddSubPage = AddSubPageString(j['Name'], x['ParentId'], j['SpgId'], "", 0)

                    $('#accordionExample' + j['ParentId']).children('.accordion-item').append(AddSubPage)

                }

            }

        }

        /**this Group */
        if (x['GroupId'] !== undefined && x['GroupId'] != 0 && x['NewGroupId'] == 0 && x['PgId'] === undefined) {

            var AddGroup1 = AddGroupString(x['Name'], x['GroupId'], 0)

            $('.newpages').append(AddGroup1)

            for (let y of overallarray) {

                if ((x['GroupId'] == y['Pgroupid']) && y['GroupId'] === undefined && y['NewGrpId'] == 0) {

                    var AddPage = AddPageString(y['Name'], y['PgId'], y['NewPgId'])

                    $('.groupdiv' + x['GroupId']).append(AddPage)

                } else if ((x['GroupId'] == y['NewGrpId']) && y['GroupId'] === undefined && y['Pgroupid'] == 0) {

                    var AddPage = AddPageString(y['Name'], y['PgId'], y['NewPgId'])

                    $('.groupdiv' + x['GroupId']).append(AddPage)

                }

            }

        } else if (x['GroupId'] !== undefined && x['GroupId'] == 0 && x['NewGroupId'] != 0 && x['PgId'] === undefined) {

            var AddGroup1 = AddGroupString(x['Name'], x['NewGroupId'], 0)

            $('.newpages').append(AddGroup1)

            for (let y of overallarray) {

                if ((x['NewGroupId'] == y['Pgroupid']) && y['GroupId'] === undefined && y['NewGrpId'] == 0) {

                    var AddPage = AddPageString(y['Name'], y['PgId'], 0)

                    $('.groupdiv' + x['NewGroupId']).append(AddPage)

                } else if ((x['NewGroupId'] == y['NewGrpId']) && y['GroupId'] === undefined && y['Pgroupid'] == 0) {

                    var AddPage = AddPageString(y['Name'], 0, y['NewPgId'])

                    $('.groupdiv' + x['NewGroupId']).append(AddPage)

                }

            }
        }


    }

    for (let x of overallarray) {

        /**this sub */
        for (let j of Subpage) {

            if ((j['NewParentId'] == x['PgId'] || j['NewParentId'] == x['NewPgId']) && j['ParentId'] == 0) {

                suborderindex = j['OrderIndex']

                var AddSubPage = AddSubPageString(j['Name'], j['ParentId'] == 0 ? j['NewParentId'] : j['ParentId'], j['SpgId'], "", j['NewParentId'])

                $('#accordionExample' + j['NewParentId']).children('.accordion-item').children('.accordion-header').children('a').attr('data-bs-target', '#collapse' + j['NewParentId']).attr('aria-controls', 'collapse' + j['NewParentId'])

                $(`<img src="/public/img/caret-down-copy-21.svg" style="position: absolute;right: 10px;">`).insertBefore($('.subpg' + j['NewParentId']))

                $('#accordionExample' + j['NewParentId']).children('.accordion-item').append(AddSubPage)

            } else if ((j['ParentId'] == x['PgId'] || ['ParentId'] == x['NewPgId']) && j['NewParentId'] == 0) {

                suborderindex = j['OrderIndex']

                var AddSubPage = AddSubPageString(j['Name'], j['ParentId'] == 0 ? j['NewParentId'] : j['ParentId'], j['SpgId'], "", j['NewParentId'])

                $('#accordionExample' + j['ParentId']).children('.accordion-item').children('.accordion-header').children('a').attr('data-bs-target', '#collapse' + j['ParentId']).attr('aria-controls', 'collapse' + j['ParentId'])

                $(`<img src="/public/img/caret-down-copy-21.svg" style="position: absolute;right: 10px;">`).insertBefore($('.subpg' + j['ParentId']))

                $('#accordionExample' + j['ParentId']).children('.accordion-item').append(AddSubPage)

            }

        }

    }

    suborderindex = suborderindex + 1

    $('.accordion-collapse').removeClass('show')

    $('.page').each(function () {
        $(this).trigger('click')
        return false
    })

}

/** Revision  */
$(document).on('click', '.revision-btn', function () {

    $('.revisionmenu').show();

})

/** Resivion Close */
$(document).on('click', '.revisionclose', function () {

    $('.revisionmenu').hide();
})


// Savebtn
$(document).on('click', '#savebtn', function () {
    
    var spid = $("#spid").val();

    $.ajax({
        url: "/spaces/page/pagecreate",
        type: "POST",
        dataType: "json",
        data: {
            "spaceid": spid,
            "creategroups": JSON.stringify({ newGrouparr }),
            "createpages": JSON.stringify({ newpagesarr }),
            "createsubpage":JSON.stringify({Subpagearr }),
            "deletepage":JSON.stringify({deletePagearr }),
            "deletegroup":JSON.stringify({deletegrouparr}),
            "deletesub" : JSON.stringify({deletesubarr}),
            "save": "save",
            csrf: $("input[name='csrf']").val()
        },
        success: function (result) {
            if (result){

                setCookie('get-toast','Pages Updated Successfully',1)
                setCookie('Alert-msg','success',1)

            } else {
           
                setCookie('Alert-msg','Internal Server Error',1)
            }

            // window.location.reload(true);
            
        }
    })
})