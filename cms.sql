--Insert Default Values


INSERT INTO tbl_roles(id, name, description, slug, is_active, is_deleted, created_on, created_by) VALUES (1, 'Super Admin', 'Default user type. Has the full administration power', 'super_admin', 1, 0, 'current-time', 1);

INSERT INTO tbl_languages(id,language_name,language_code,image_path,json_path,is_status,is_default,created_by,created_on,is_deleted) VALUES (1,'English', 'en', '/public/img/in.jpeg','locales/en.json', 1, 1,1, 'current-time',0)

INSERT INTO tbl_languages(id,language_name,language_code,image_path,json_path,is_status,is_default,created_by,created_on,is_deleted) VALUES (2,'Spanish', 'es','/public/img/sp.jpeg', 'locales/es.json', 1, 0,1, 'current-time',0)

INSERT INTO tbl_languages(id,language_name,language_code,image_path,json_path,is_status,is_default,created_by,created_on,is_deleted) VALUES (3,'French', 'fr','/public/img/fr.jpeg', 'locales/fr.json', 1, 0,1, 'current-time',0)

INSERT INTO tbl_languages(id,language_name,language_code,image_path,json_path,is_status,is_default,created_by,created_on,is_deleted) VALUES (4,'Russian', 'ru','/public/img/russia.svg', 'locales/ru.json', 1, 0,1, 'current-time',0)

INSERT INTO tbl_user_personalizes(id,user_id,menu_background_color,logo_path,expand_logo_path) values(1,1,'rgb(9, 171, 217)','/public/img/logo.svg','/public/img/logo-bg.svg');

-- INSERT INTO tbl_storage_types(id, local, selected_type, aws) VALUES (1, 'storage', 'aws','awsConfig');


INSERT INTO tbl_storage_types(id, local, selected_type) VALUES (1, 'storage', 'aws');


--Channel default fields

INSERT INTO tbl_field_types(id, type_name, type_slug, is_active, is_deleted, created_by, created_on) VALUES (1, 'Label', 'label', 1,  0, 1, 'current-time');

INSERT INTO tbl_field_types(id, type_name, type_slug, is_active, is_deleted, created_by, created_on) VALUES (2, 'Text', 'text', 1,  0, 1, 'current-time');

INSERT INTO tbl_field_types(id, type_name, type_slug, is_active, is_deleted, created_by, created_on) VALUES (3, 'Link', 'link', 1,  0, 1, 'current-time');

INSERT INTO tbl_field_types(id, type_name, type_slug, is_active, is_deleted, created_by, created_on) VALUES (4, 'Date & Time', 'date&time', 1,  0, 1, 'current-time');

INSERT INTO tbl_field_types(id, type_name, type_slug, is_active, is_deleted, created_by, created_on) VALUES (5, 'Select', 'select', 1,  0, 1, 'current-time');

INSERT INTO tbl_field_types(id, type_name, type_slug, is_active, is_deleted, created_by, created_on) VALUES (6, 'Date', 'date', 1,  0, 1, 'current-time');

INSERT INTO tbl_field_types(id, type_name, type_slug, is_active, is_deleted, created_by, created_on) VALUES (7, 'TextBox', 'textbox', 1,  0, 1, 'current-time');

INSERT INTO tbl_field_types(id, type_name, type_slug, is_active, is_deleted, created_by, created_on) VALUES (8, 'TextArea', 'textarea', 1, 0, 1, 'current-time');

INSERT INTO tbl_field_types(id, type_name, type_slug, is_active, is_deleted, created_by, created_on) VALUES (9, 'Radio Button', 'radiobutton', 1, 0, 1, 'current-time');

INSERT INTO tbl_field_types(id, type_name, type_slug, is_active, is_deleted, created_by, created_on) VALUES (10, 'CheckBox', 'checkbox', 1, 0, 1, 'current-time');

INSERT INTO tbl_field_types(id, type_name, type_slug, is_active, is_deleted, created_by, created_on) VALUES (11, 'Text Editor', 'texteditor', 1, 0, 1, 'current-time');

INSERT INTO tbl_field_types(id, type_name, type_slug, is_active, is_deleted, created_by, created_on) VALUES (12, 'Section', 'section', 1, 0, 1, 'current-time');

INSERT INTO tbl_field_types(id, type_name, type_slug, is_active, is_deleted, created_by, created_on) VALUES (13, 'Section Break', 'sectionbreak', 1, 0, 1, 'current-time');

INSERT INTO tbl_field_types(id, type_name, type_slug, is_active, is_deleted, created_by, created_on) VALUES (14, 'Members', 'member', 1,  0, 1, 'current-time');

INSERT INTO tbl_field_types(id, type_name, type_slug, is_active, is_deleted, created_by, created_on) VALUES (15, 'Media Gallery', 'mediagallrey', 1,  0, 1, 'current-time');

INSERT INTO tbl_field_types(id, type_name, type_slug, is_active, is_deleted, created_by, created_on) VALUES (16, 'Video URL', 'videourl', 1,  0, 1, 'current-time');


--Default Insert Menu value

INSERT INTO tbl_modules(id, module_name, is_active, created_by, created_on, default_module, parent_id, assign_permission, icon_path, description, order_index, menu_type,full_access_permission,group_flg) VALUES (1, 'Dashboard', 1, 1, 'current-time', 0, 0, 1, '/public/img/dashboard_new.svg', '', 1, 'left',0,0);

INSERT INTO tbl_modules(id, module_name, is_active, created_by, created_on, default_module, parent_id, assign_permission, icon_path, description, order_index, menu_type,full_access_permission,group_flg) VALUES (2, 'Categories', 1, 1, 'current-time', 0, 0, 1, '/public/img/category_new.svg', 'Create various categories for content and organize them efficiently by categorizing it with user-defined labels.', 2, 'left',0,0);

INSERT INTO tbl_modules(id, module_name, is_active, created_by, created_on, default_module, parent_id, assign_permission, icon_path, description, order_index, menu_type,full_access_permission,group_flg) VALUES(3, 'Channels', 1, 1, 'current-time', 0, 0, 1, '/public/img/cms-on.svg', 'Easily create and manage diverse content types using Spurt CMS Admin.', 3, 'left',0,0);

INSERT INTO tbl_modules(id, module_name, is_active, created_by, created_on, default_module, parent_id, assign_permission, icon_path, description, order_index, menu_type,full_access_permission,group_flg) VALUES(4, 'Entries', 1, 1, 'current-time', 0, 0, 1, '/public/img/left-entries.svg', 'Seamlessly add and update content within specific channels for targeted delivery and right audience. ', 4, 'tab',0,0);

INSERT INTO tbl_modules(id, module_name, is_active, created_by, created_on, default_module, parent_id, assign_permission, icon_path, description, order_index, menu_type,full_access_permission,group_flg) VALUES(5, 'Members', 1, 1, 'current-time', 0, 0, 1, '/public/img/member_new.svg', 'Admin manages member profiles, including adding new members and assigning them to specific groups.', 5, 'left',0,0);

INSERT INTO tbl_modules(id, module_name, is_active, created_by, created_on, default_module, parent_id, assign_permission, icon_path, description, order_index, menu_type,full_access_permission,group_flg) VALUES(6, 'Settings', 1, 1, 'current-time', 0, 0, 1, '/public/img/settings.svg', 'Check relevant boxes for providing access control to this role, in this settings.', 6, 'left',0,1);

INSERT INTO tbl_modules(id, module_name, is_active, created_by, created_on, default_module, parent_id, assign_permission, icon_path, description, order_index, menu_type,full_access_permission,group_flg) VALUES(7, 'Channels', 1, 1, 'current-time', 0, 3, 0, '/public/img/accord-channels.svg', 'Easily create and manage diverse content types using Spurt CMS Admin.', 7, 'tab',1,0);

INSERT INTO tbl_modules(id, module_name, is_active, created_by, created_on, default_module, parent_id, assign_permission, icon_path, description, order_index, menu_type,full_access_permission,group_flg) VALUES(8, 'Entries', 1, 1, 'current-time', 0, 4, 0, '/public/img/entries.svg', 'Seamlessly add and update content within specific channels for targeted delivery and right audience. ', 8, 'tab',1,0);

INSERT INTO tbl_modules(id, module_name, is_active, created_by, created_on, default_module, parent_id, assign_permission, icon_path, description, order_index, menu_type,full_access_permission,group_flg) VALUES(9, 'Categories Group', 1, 1, 'current-time', 0, 2, 0, '/public/img/categories-group.svg', 'Create various categories for content and organize them efficiently by categorizing it with user-defined labels.', 9, 'tab',0,0);

INSERT INTO tbl_modules(id, module_name, is_active, created_by, created_on, default_module, parent_id, assign_permission, icon_path, description, order_index, menu_type,full_access_permission,group_flg) VALUES(10, 'Categories', 1, 1, 'current-time', 0, 2, 0, '/public/img/categories-group.svg', '', 10, 'tab',0,0);

INSERT INTO tbl_modules(id, module_name, is_active, created_by, created_on, default_module, parent_id, assign_permission, icon_path, description, order_index, menu_type,full_access_permission,group_flg) VALUES(11, 'Member', 1, 1, 'current-time', 0, 5, 0, '/public/img/accord-member.svg', 'Add member profiles, map them to a member group and manage the entire list of member profiles. ', 11, 'tab',0,0);

INSERT INTO tbl_modules(id, module_name, is_active, created_by, created_on, default_module, parent_id, assign_permission, icon_path, description, order_index, menu_type,full_access_permission,group_flg) VALUES(12, 'Members Group', 1, 1, 'current-time', 0, 5, 0, '/public/img/group.svg', 'Create groups and categorize members into various groups like Elite Members or Favorite Members', 12, 'tab',0,0);

INSERT INTO tbl_modules(id, module_name, is_active, created_by, created_on, default_module, parent_id, assign_permission, icon_path, description, order_index, menu_type,full_access_permission,group_flg) VALUES(13, 'Member Restrict', 1, 1, 'current-time', 0, 5, 0, '/public/img/accord-memberRestrict.svg', '', 13, 'tab',1,0);

INSERT INTO tbl_modules(id, module_name, is_active, created_by, created_on, default_module, parent_id, assign_permission, icon_path, description, order_index, menu_type,full_access_permission,group_flg) VALUES(14, 'My Account', 1, 1, 'current-time', 0, 6, 1, '/public/img/profile.svg', 'Here you can change your email id and password.', 14, 'tab',1,0)

INSERT INTO tbl_modules(id, module_name, is_active, created_by, created_on, default_module, parent_id, assign_permission, icon_path, description, order_index, menu_type,full_access_permission,group_flg) VALUES(15, 'Personalize', 1, 1, 'current-time', 0, 6, 1, '/public/img/personalize.svg', 'You can change the theme according your preference. ', 15, 'tab',1,0)

INSERT INTO tbl_modules(id, module_name, is_active, created_by, created_on, default_module, parent_id, assign_permission, icon_path, description, order_index, menu_type,full_access_permission,group_flg) VALUES(16, 'Roles & Permissions', 1, 1, 'current-time', 0, 6, 0, '/public/img/roles-permision.svg', 'Create different admin roles and assign permissions to them based on their responsibilities.', 16, 'tab',1,0)

INSERT INTO tbl_modules(id, module_name, is_active, created_by, created_on, default_module, parent_id, assign_permission, icon_path, description, order_index, menu_type,full_access_permission,group_flg) VALUES(17, 'Team', 1, 1, 'current-time', 0, 6, 0, '/public/img/team.svg', 'Create a team of different admin roles.', 17, 'tab',1,0)

INSERT INTO tbl_modules(id, module_name, is_active, created_by, created_on, default_module, parent_id, assign_permission, icon_path, description, order_index, menu_type,full_access_permission,group_flg) VALUES(18, 'Email', 1, 1, 'current-time', 0, 6, 0, '/public/img/email.svg', 'Create and Maintain Email Templates that are to be sent to users in different scenarios.', 18, 'tab',1,0)

INSERT INTO tbl_modules(id, module_name, is_active, created_by, created_on, default_module, parent_id, assign_permission, icon_path, description, order_index, menu_type,full_access_permission,group_flg) VALUES(19, 'Dashboard', 1, 1, 'current-time', 0, 1, 1, '/public/img/settings.svg', '', 19, 'tab',1,0);

INSERT INTO tbl_modules(id, module_name, is_active, created_by, created_on, default_module, parent_id, assign_permission, icon_path, description, order_index, menu_type,full_access_permission,group_flg) VALUES(20, 'Settings', 1, 1, 'current-time', 0, 6, 1, '/public/img/settings.svg', '', 20, 'tab',1,0)

INSERT INTO tbl_modules(id, module_name, is_active, created_by, created_on, default_module, parent_id, assign_permission, icon_path, description, order_index, menu_type,full_access_permission,group_flg) VALUES(21, 'Settings', 1, 1, 'current-time', 0, 3, 1, '/public/img/settings.svg', '', 21, 'tab',1,0)

INSERT INTO tbl_modules(id, module_name, is_active, created_by, created_on, default_module, parent_id, assign_permission, icon_path, description, order_index, menu_type,full_access_permission,group_flg) VALUES(22, 'Settings', 1, 1, 'current-time', 0, 5, 1, '/public/img/settings.svg', '', 22, 'tab',1,0)

INSERT INTO tbl_modules(id, module_name, is_active, created_by, created_on, default_module, parent_id, assign_permission, icon_path, description, order_index, menu_type,full_access_permission,group_flg) VALUES(23, 'General Settings', 1, 1, 'current-time', 0, 6, 0, '/public/img/settings.svg', '', 23, 'tab',1,0)

INSERT INTO tbl_modules(id, module_name, is_active, created_by, created_on, default_module, parent_id, assign_permission, icon_path, description, order_index, menu_type,full_access_permission,group_flg) VALUES(24, 'Graphql Playground', 1, 1, 'current-time', 0, 0, 0, '/public/img/graphql-playground.svg', '', 24, 'left',1,1)

INSERT INTO tbl_modules(id, module_name, is_active, created_by, created_on, default_module, parent_id,assign_permission, icon_path, description, order_index, menu_type,full_access_permission,group_flg) VALUES(25, 'Graphql Playground', 1, 1, 'current-time', 0, 24, 0, '/public/img/graphql-playground.svg', '', 25, 'tab',1,0)

INSERT INTO tbl_modules(id, module_name, is_active, created_by, created_on, default_module, parent_id, assign_permission, icon_path, description, order_index, menu_type,full_access_permission,group_flg) VALUES(26, 'Template', 1, 1, 'current-time', 0, 0, 0, '/public/img/templates.svg', '', 26, 'left',1,1)

INSERT INTO tbl_modules(id, module_name, is_active, created_by, created_on, default_module, parent_id, assign_permission, icon_path, description, order_index, menu_type,full_access_permission,group_flg) VALUES(27,  'Template', 1, 1, 'current-time', 0, 26, 0, '/public/img/templates.svg', '', 27, 'tab',1,0)

INSERT INTO tbl_modules(id, module_name, is_active, created_by, created_on, default_module, parent_id, assign_permission, icon_path, description, order_index, menu_type,full_access_permission,group_flg) VALUES(28, 'Graphql Api', 1, 1, 'current-time', 0, 0, 0, '/public/img/role-graph-ql.svg', '', 28, 'left',1,1)

INSERT INTO tbl_modules(id, module_name, is_active, created_by, created_on, default_module, parent_id, assign_permission, icon_path, description, order_index, menu_type,full_access_permission,group_flg) VALUES(29, 'Graphql Api', 1, 1, 'current-time', 0, 28, 0, '/public/img/Graph-QL.svg', '', 29, 'tab',1,0)

INSERT INTO tbl_modules(id, module_name, is_active, created_by, created_on, default_module, parent_id, assign_permission, icon_path, description, order_index, menu_type,full_access_permission,group_flg) VALUES(30, 'Languages', 1, 1, 'current-time', 0, 0, 0, '/public/img/language.svg', '', 30, 'left',1,1)

INSERT INTO tbl_modules(id, module_name, is_active, created_by, created_on, default_module, parent_id, assign_permission, icon_path, description, order_index, menu_type,full_access_permission,group_flg) VALUES(31, 'Languages', 1, 1, 'current-time', 0, 30, 0, '/public/img/language.svg', '', 31, 'tab',1,0)


--Default Module Permission Routes

INSERT INTO tbl_module_permissions(id, route_name, display_name, description, module_id, created_by, created_on, full_access_permission, parent_id, assign_permission,order_index, slug_name) VALUES (1, '/dashboard', 'Dashboard', '', 19, 1, 'current-time', 1, 0, 1, 1, 'dashboard')

INSERT INTO tbl_module_permissions(id, route_name, display_name, description, module_id, created_by, created_on, full_access_permission, parent_id, assign_permission,order_index, slug_name) VALUES (2, '/categories/', 'View', 'Give view access to the category group', 9, 1, 'current-time', 0, 0, 1, 1, 'view')

INSERT INTO tbl_module_permissions(id, route_name, display_name, description, module_id, created_by, created_on, full_access_permission, parent_id, assign_permission,order_index, slug_name) VALUES (3, '/categories/newcategory', 'Create', 'Give create access to the category Group', 9, 1, 'current-time', 0, 0, 1, 2, 'create')

INSERT INTO tbl_module_permissions(id, route_name, display_name, description, module_id, created_by, created_on, full_access_permission, parent_id, assign_permission,order_index, slug_name) VALUES (4, '/categories/updatecategory', 'Update', 'Give update access to the category Group', 9, 1, 'current-time', 0, 0, 1, 3, 'update')

INSERT INTO tbl_module_permissions(id, route_name, display_name, description, module_id, created_by, created_on, full_access_permission, parent_id, assign_permission,order_index, slug_name) VALUES (5, '/categories/deletecategory', 'Delete', 'Give delete access to the category Group', 9, 1, 'current-time', 0, 0, 1, 4, 'delete')

INSERT INTO tbl_module_permissions(id, route_name, display_name, description, module_id, created_by, created_on, full_access_permission, parent_id, assign_permission,order_index, slug_name) VALUES (6, '/categories/addcategory/','View','Give view access to the categories', 10, 1, 'current-time', 0, 0, 1, 1, 'view')

INSERT INTO tbl_module_permissions(id, route_name, display_name, description, module_id, created_by, created_on, full_access_permission, parent_id, assign_permission,order_index, slug_name) VALUES (7, '/categories/addsubcategory', 'Create', 'Give create access to the categories', 10, 1, 'current-time', 0, 0, 1, 2, 'create')

INSERT INTO tbl_module_permissions(id, route_name, display_name, description, module_id, created_by, created_on, full_access_permission, parent_id, assign_permission,order_index, slug_name) VALUES (8, '/categories/editsubcategory', 'Update', 'Give update access to the categories', 10, 1, 'current-time', 0, 0, 1, 3, 'update')

INSERT INTO tbl_module_permissions(id, route_name, display_name, description, module_id, created_by, created_on, full_access_permission, parent_id, assign_permission,order_index, slug_name) VALUES (9, '/categories/removecategory', 'Delete', 'Give delete access to the categories', 10, 1, 'current-time', 0, 0, 1, 4, 'delete')

INSERT INTO tbl_module_permissions(id, route_name, display_name, description, module_id, created_by, created_on, full_access_permission, parent_id, assign_permission,order_index, slug_name) VALUES (10, '/channels/', 'Channel', 'Give full access to the channels', 7, 1, 'current-time', 1, 0, 1, 1, 'channels')

INSERT INTO tbl_module_permissions(id, route_name, display_name, description, module_id, created_by, created_on, full_access_permission, parent_id, assign_permission,order_index, slug_name) VALUES (11, '/membersgroup/', 'View', 'Give view access to the member group', 12, 1, 'current-time', 0, 0,1, 1, 'view')

INSERT INTO tbl_module_permissions(id, route_name, display_name, description, module_id, created_by, created_on, full_access_permission, parent_id, assign_permission,order_index, slug_name) VALUES (12, '/membersgroup/newgroup', 'Create', 'Give create access to the member group', 12, 1, 'current-time', 0, 0, 1, 2, 'create')

INSERT INTO tbl_module_permissions(id, route_name, display_name, description, module_id, created_by, created_on, full_access_permission, parent_id, assign_permission,order_index, slug_name) VALUES (13, '/membersgroup/updategroup', 'Update', 'Give update access to the member group', 12, 1, 'current-time', 0, 0, 1, 3, 'update')

INSERT INTO tbl_module_permissions(id, route_name, display_name, description, module_id, created_by, created_on, full_access_permission, parent_id, assign_permission,order_index, slug_name) VALUES (14, '/membersgroup/deletegroup', 'Delete', 'Give delete access to the member group', 12, 1, 'current-time', 0, 0, 1, 4, 'delete')

INSERT INTO tbl_module_permissions(id, route_name, display_name, description, module_id, created_by, created_on, full_access_permission, parent_id, assign_permission,order_index, slug_name) VALUES (15, '/member/', 'View', 'Give view access to the Member', 11, 1, 'current-time', 0, 0, 1, 1, 'view')

INSERT INTO tbl_module_permissions(id, route_name, display_name, description, module_id, created_by, created_on, full_access_permission, parent_id, assign_permission,order_index, slug_name) VALUES (16, '/member/newmember', 'Create', 'Give create access to the Member', 11, 1, 'current-time', 0, 0, 1, 2, 'create')

INSERT INTO tbl_module_permissions(id, route_name, display_name, description, module_id, created_by, created_on, full_access_permission, parent_id, assign_permission,order_index, slug_name) VALUES (17, '/member/updatemember', 'Update', 'Give update access to the Member', 11, 1, 'current-time', 0, 0, 1, 3, 'update')

INSERT INTO tbl_module_permissions(id, route_name, display_name, description, module_id, created_by, created_on, full_access_permission, parent_id, assign_permission,order_index, slug_name) VALUES (18, '/member/deletemember', 'Delete', 'Give delete access to the Member', 11, 1, 'current-time', 0, 0, 1, 4, 'delete')

INSERT INTO tbl_module_permissions(id, route_name, display_name, description, module_id, created_by, created_on, full_access_permission, parent_id, assign_permission,order_index, slug_name) VALUES (19, '/memberaccess/', 'Member-Restrict', 'Create and manage content consumption access and restrictions to the Members on the website. ', 13, 1, 'current-time', 1, 0, 1, 1, 'memberaccess')

INSERT INTO tbl_module_permissions(id, route_name, display_name, description, module_id, created_by, created_on, full_access_permission, parent_id, assign_permission,order_index, slug_name) VALUES (20, '/channel/entrylist', 'Entries', '', 8, 1, 'current-time', 1, 0, 0, 0, 'entries')

INSERT INTO tbl_module_permissions(id, route_name, display_name, description, module_id, created_by, created_on, full_access_permission, parent_id, assign_permission,order_index, slug_name) VALUES (21, '/settings/roles/', 'Roles & Permission', 'Create and manage additional roles based on the needs for their tasks and activities in the control panel.', 16, 1, 'current-time', 1, 0, 1, 1, 'roles')

INSERT INTO tbl_module_permissions(id, route_name, display_name, description, module_id, created_by, created_on, full_access_permission, parent_id, assign_permission,order_index, slug_name) VALUES (22, '/settings/users/userlist', 'Teams', 'Add users for assigning them roles, followed by assigning permissions for appropriate access to the CMS.', 17, 1, 'current-time', 1, 0, 1, 1, 'users')

INSERT INTO tbl_module_permissions(id, route_name, display_name, description, module_id, created_by, created_on, full_access_permission, parent_id, assign_permission,order_index, slug_name) VALUES (23, '/settings/emails', 'Email Templates', 'Add different languages here, for the Members to consume content in the language of preference on the website.', 18, 1, 'current-time', 1, 0, 1, 1, 'email')

INSERT INTO tbl_module_permissions(id, route_name, display_name, description, module_id, created_by, created_on, full_access_permission, parent_id, assign_permission,order_index, slug_name) VALUES (24, '/member/settings', 'Settings', '', 22, 1, 'current-time', 1, 0, 0, 1, 'settings')

INSERT INTO tbl_module_permissions(id, route_name, display_name, description, module_id, created_by, created_on, full_access_permission, parent_id, assign_permission,order_index, slug_name) VALUES (25, '/settings/general-settings/', 'General Settings', 'Add different languages here, for the Members to consume content in the language of preference on the website.', 23, 1, 'current-time', 1, 0, 1, 1, 'data')

INSERT INTO tbl_module_permissions(id, route_name, display_name, description, module_id, created_by, created_on, full_access_permission, parent_id, assign_permission,order_index, slug_name) VALUES (26, '/channel/entrylist/1', 'Default Channel', '', 8, 1, 'current-time', 1, 0, 1, 1, 'Default Channel')

INSERT INTO tbl_module_permissions(id, route_name, display_name, description, module_id, created_by,created_on, full_access_permission, parent_id, assign_permission,order_index, slug_name) VALUES (27, '/graphql', 'Graphql Api', 'Give full access to the graphql api', 29, 1, 'current-time', 1, 0, 1, 1, 'graphql')

INSERT INTO tbl_module_permissions(id, route_name, display_name, description, module_id, created_by, created_on, full_access_permission, parent_id, assign_permission,order_index, slug_name) VALUES(28, '/templates', 'Template', 'Give full access to the template', 27, 1, 'current-time', 1, 0, 1, 2, 'template')

INSERT INTO tbl_module_permissions(id, route_name, display_name, description, module_id, created_by, created_on, full_access_permission, parent_id, assign_permission,order_index, slug_name) VALUES (29, '/graphql/playground', 'Graphql Playground', 'Give full access to the graphql playground', 25, 1, 'current-time', 1, 0, 1, 2, 'graphqlplayground')

INSERT INTO tbl_module_permissions(id, route_name, display_name, description, module_id, created_by, created_on, full_access_permission, parent_id, assign_permission,order_index, slug_name) VALUES (30, '/languages', 'Languages', 'Give full access to the languages', 31, 1, 'current-time', 1, 0, 1, 2, 'languages')

INSERT INTO tbl_timezones(id,timezone) VALUES (1,'Africa/Cairo'),(2,'Africa/Johannesburg'),(3,'Africa/Lagos'),(4,'Africa/Nairobi'),(5,'America/Argentina/Buenos_Aires'),(6,'America/Chicago'),(7,'America/Denver'),(8,'America/Los_Angeles'),(9,'America/Mexico_City'),(10,'America/New_York'),(11,'America/Sao_Paulo'),(12,'Asia/Bangkok'),(13,'Asia/Dhaka'),(14,'Asia/Dubai'),(15,'Asia/Hong_Kong'),(16,'Asia/Jakarta'),(17,'Asia/Kolkata'),(18,'Asia/Manila'),(19,'Asia/Seoul'),(20,'Asia/Shanghai'),(21,'Asia/Singapore'),(22,'Asia/Tokyo'),(23,'Australia/Melbourne'),(24,'Australia/Sydney'),(25,'Europe/Amsterdam'),(26,'Europe/Berlin'),(27,'Europe/Istanbul'),(28,'Europe/London'),(29,'Europe/Madrid'),(30,'Europe/Moscow'),(31,'Europe/Paris'),(32,'Europe/Rome'),(33,'Pacific/Auckland'),(34,'Pacific/Honolulu')

INSERT INTO tbl_users(id, role_id, first_name, username, password,  is_active, created_on, created_by, is_deleted, default_language_id,tenant_id,s3_folder_name)VALUES (1, 2, 'spurtCMSAdmin',  'spurtcmsAdmin', '$2a$14$3kdNuP2Fo/SBopGK9e/9ReBa8uq82vMM5Ko2jnbLpuPb6qxqgR0x2', 1, 'current-time',1, 0, 1,1,'SpurtCMS_1/');

INSERT INTO tbl_roles(id, name, description, slug, is_active, is_deleted, created_on, created_by,tenant_id) VALUES (2, 'Admin', 'Default user type. Has the full administration power', 'admin', 1, 0, 'current-time', 1,1);

INSERT INTO tbl_mstr_tenants(id,tenant_id,s3_storage_path,is_deleted)VALUES(1,1,'SpurtCMS_1/',0)

-- Must use separate insert queries for parent and child data insertions to maintain sequence.

-- Some reserved words are used for dynamic values insertion in queries.They are given as below.Use them whenever they needed essentially.

-- current-time,2(user id),2(tenant id),pid(parent id),chid(channel id),blid(block id),tagid(tag id),mapcat(mapped categories)

INSERT INTO tbl_categories(id,category_name, category_slug, created_on, created_by,is_deleted, parent_id, description, tenant_id)	VALUES (1,'Default Category', 'default_category', 'current-time', 1, 0, 0, 'Default_Category', 1);

INSERT INTO tbl_categories(id,category_name, category_slug, created_on, created_by,is_deleted, parent_id, description, tenant_id)	VALUES (2,'Default1', 'default1', 'current-time', 1, 0, 1, 'Default1',1);

INSERT INTO tbl_channels(id,channel_name, slug_name, field_group_id, is_active, is_deleted, created_on, created_by,channel_description,channel_type,tenant_id) VALUES (1,'Default_Channel', 'default_channel', 0, 1, 0, 'current-time', 1, 'default description','mychannels',1);

INSERT INTO tbl_channel_categories(id,channel_id, category_id, created_at, created_on,tenant_id) VALUES (1,1,'1,2', 1, 'current-time',1);

INSERT INTO tbl_member_groups(id,name,slug,description,is_active,is_deleted,created_on,created_by,tenant_id) VALUES (1,'Default Group', 'default-group', 'Default Group', 1,0, 'current-time', 1,1);

INSERT INTO tbl_member_settings (id,allow_registration,member_login,notification_users,tenant_id) VALUES (1,1,'password','1',1)

INSERT INTO tbl_email_configurations (id,selected_type,tenant_id) VALUES (1,'environment',1)

INSERT INTO tbl_general_settings (id,date_format,expand_logo_path,language_id,logo_path,storage_type,tenant_id,time_format,time_zone) VALUES (2,'dd mmm yyyy','/public/img/logo-bg.svg',1,'/public/img/logo1.svg','aws',1,'12','Asia/Kolkata')

INSERT INTO tbl_graphql_settings (id,token_name,description,duration,created_by,created_on,is_deleted,token,is_default,tenant_id) VALUES (1,'Default Token','Default Token','Unlimited',1,'current-time',0,'apikey',1,1)

-- below insert queries are for CREATING DEFAULT CATEGORIES, CHANNELS and TEMPLATES 

-- ****************** Knowledge-Base ***********************
--  default knowledge base categories for channel based templates
INSERT INTO tbl_categories(id,category_name, category_slug, created_on, created_by, is_deleted, parent_id, description, tenant_id) VALUES (3,'Knowledge Base', 'knowledge-base', 'current-time', 1, 0,0,'Parent Category for knowledge base', 1)

-- default knowledge base channels for channel based templates
INSERT INTO tbl_channels(id,channel_name, slug_name, field_group_id, is_active, is_deleted, created_on, created_by,channel_description,channel_type,tenant_id) VALUES (2,'Knowledge Base', 'knowledge-base', 0, 1, 0, 'current-time', 1, 'Channel for knowledge base templates','mychannels',1);

INSERT INTO tbl_module_permissions(id,route_name,display_name,slug_name,module_id,full_access_permission,parent_id,assign_permission,order_index,created_by,created_on,tenant_id) VALUES (31,'/channel/entrylist/2','Knowledge Base','knowledge-base',8,1,0,1,2,1,'current-time',1)

-- INSERT INTO tbl_channel_categories(channel_id, category_id, created_at, created_on,tenant_id) VALUES (chid,'mapcat', 2, 'current-time',1);

-- default knowledge-base template module creation
INSERT INTO tbl_template_modules(id,template_module_name,template_module_slug,description,is_active,created_by,created_on,channel_id,tenant_id) VALUES (1,'Help Center', 'help-center', 'Help Center Templates', 1, 1, 'current-time', 2,1)

-- default knowledge-base template creation
INSERT INTO tbl_templates(id,template_name, template_slug, template_module_id, image_name, image_path, deploy_link, github_link, created_on, created_by, is_active, is_deleted,channel_id,tenant_id) VALUES(1,'Knowledge Base Theme','knowledge-base-theme', 1,'knowledgeBase.png','media/knowledgeBase.png','https://vercel.com/new/clone?repository-url=https://github.com/spurtcms/knowledge-base&demo-title=knowledge-base&env=NEXT_PUBLIC_SPURTCMS_NEXTJS_STARTER_APIKEY','https://github.com/spurtcms/knowledge-base','current-time',1,1,0,2,1)


-- ***************************** Blogs *********************************
-- Please DON'T change the flow of the input statements it will definitely affect the output.

--  default Parent blog CATEGORY for channel based templates
INSERT INTO tbl_categories(id,category_name, category_slug, created_on, created_by, is_deleted, parent_id, description, tenant_id) VALUES (4,'Blog', 'blog', 'current-time', 1, 0,0,'Parent Category for blogs', 1)

-- default Parent blog CHANNEL for channel based templates
INSERT INTO tbl_channels(id,channel_name, slug_name, field_group_id, is_active, is_deleted, created_on, created_by,channel_description,channel_type,tenant_id) VALUES (4,'Blog', 'blog', 0, 1, 0, 'current-time', 1, 'Channel for blog templates','mychannels',1);

INSERT INTO tbl_module_permissions(id,route_name,display_name,slug_name,module_id,full_access_permission,parent_id,assign_permission,order_index,created_by,created_on,tenant_id) VALUES (32,'/channel/entrylist/3','Blog','blog',8,1,0,1,2,1,'current-time',1)


-- default blog template module creation
INSERT INTO tbl_template_modules(id,template_module_name,template_module_slug,description,is_active,created_by,created_on,channel_id,tenant_id) VALUES (2,'Blog', 'blog', 'Blog Templates', 1, 1, 'current-time',3,1)

-- default NextJS Starter Theme template creation
INSERT INTO tbl_templates(id,template_name, template_slug, template_module_id, image_name, image_path, deploy_link, github_link, created_on, created_by, is_active, is_deleted,channel_id,tenant_id) VALUES(2,'NextJS Starter Theme','next-js-starter-theme', 2,'blogstarter.png','media/blogstarter.png','https://vercel.com/new/clone?repository-url=https://github.com/spurtcms/nextjs-starter-theme&demo-title=nextjs-starter-theme&env=NEXT_PUBLIC_SPURTCMS_NEXTJS_STARTER_APIKEY','https://github.com/spurtcms/nextjs-starter-theme','current-time',1,1,0,3,1)

-- default NextJS Blog Theme template creation
INSERT INTO tbl_templates(id,template_name, template_slug, template_module_id, image_name, image_path, deploy_link, github_link, created_on, created_by, is_active, is_deleted,channel_id,tenant_id) VALUES(3,'NextJS Blog Theme','next-js-blog-theme', 2,'blog1.png','media/blog1.png','https://vercel.com/new/clone?repository-url=https://github.com/spurtcms/nextjs-blog-theme&demo-title=nextjs-blog-theme&env=NEXT_PUBLIC_SPURTCMS_NEXTJS_STARTER_APIKEY','https://github.com/spurtcms/nextjs-blog-theme','current-time',1,1,0,3,1)


-- default NextJS Blogs2 Theme template creation
INSERT INTO tbl_templates(id,template_name, template_slug, template_module_id, image_name, image_path, deploy_link, github_link, created_on, created_by, is_active, is_deleted,channel_id,tenant_id) VALUES(4,'NextJS Blog2 Theme','next-js-blog2-theme', 2,'blog2.png','media/blog2.png','https://vercel.com/new/clone?repository-url=https://github.com/spurtcms/nextjs-blog2-theme&demo-title=nextjs-blog2-theme&env=NEXT_PUBLIC_SPURTCMS_NEXTJS_STARTER_APIKEY','https://github.com/spurtcms/nextjs-blog2-theme','current-time',1,1,0,3,1)


-- default NextJS Blogs4 Theme template creation
INSERT INTO tbl_templates(id,template_name, template_slug, template_module_id, image_name, image_path, deploy_link, github_link, created_on, created_by, is_active, is_deleted,channel_id,tenant_id) VALUES(5,'NextJS Blog4 Theme','next-js-blog4-theme', 2,'blog4.png','media/blog4.png','https://vercel.com/new/clone?repository-url=https://github.com/spurtcms/nextjs-blog4-theme&demo-title=nextjs-blog4-theme&env=NEXT_PUBLIC_SPURTCMS_NEXTJS_STARTER_APIKEY','https://github.com/spurtcms/nextjs-blog4-theme','current-time',1,1,0,3,1)


INSERT INTO tbl_email_templates(id, template_slug, template_subject, template_message, template_description, module_id, created_on, created_by, is_deleted, is_active,template_name,tenant_id)VALUES (1,'createuser', 'Confirm user registration', '<tr><td><p style="margin-left:0;">&nbsp;</p><h1 style="font-size: 20px; font-weight: bold;line-height: 27px;color:#000000; margin:0 0  12px 0;">Dear <strong>{FirstName}</strong>,</h1></td></tr><tr><td><p style="color:#000000;font-size:14px;">Congratulations! Your SpurtCMS account has been successfully created.</p><p style="color:#000000;font-size:14px;">Log in at</p><p  style="color:#000000;font-size:14px;">{Loginurl}</p><p  style="color:#000000;font-size:14px;margin:0 0 16px;">and start using your Admin Account.</p></td></tr><tr><td><p style="color:#000000;font-size:16px;line-height:normal;margin:0 0 12px;">Best Regards,</p><p style="color:#000000;font-size:16px;font-weight:500;line-height:24px;margin:0 0 16px;"><strong>Spurt CMS Admin</strong></p></td></tr>', 'Change Password: Enable this toggle to activate the Change Password email template, which confirms to users that their password has been successfully changed.', 0, 'current-time', 1, 0, 1,'Create Admin User',1);

INSERT INTO tbl_email_templates(id, template_slug, template_subject, template_message, template_description, module_id, created_on, created_by, is_deleted, is_active,template_name,tenant_id)VALUES (2, 'changepassword', 'New password', '<p>Hello,{firstname} your password updated Successfully. your new password:{password}</p>', 'Change Password: Change Password email template, which confirms to users that their password has been successfully changed.', 0, 'current-time', 1, 0, 1,'Change Password',1);

-- INSERT INTO tbl_email_templates(id, template_slug,template_subject,template_description,template_message,module_id,created_on,created_by,is_deleted,is_active,template_name,tenant_id)VALUES(3,'OTPGenerate','Confirm OTPs','OTP email template, which sends a One-Time Password (OTP) to users for verification.','<tr><td><h1 style="font-size: 20px; font-weight: bold;line-height: 27px;color:#000000; margin:0 0  16px 0;">Dear {FirstName}</h1></td></tr><tr><td><p style="font-size: 14px; line-height: 24px; color: #000000;margin:0;">As requsted, here is your One-Time Password (OTP) to log in to SpurtCMS:</p><h3 style="font-size: 20px; line-height:1; font-weight: bold; color: #000000; margin:20px 0;">OTP:{OTP}</h3><p style="font-size: 14px; line-height: 24px; color: #000000;margin:0 0 20px 0;">Please use this OTP within <span style="font-size: 16px; font-weight:bold;">{Expirytime}</span> to complete the login process. If you do not log in within this time frame, you may request a new OTP.</p><p style="font-size: 14px; line-height: 24px; color: #000000;margin:0 0 20px 0;">If you encounter any issue or need assistance, feel free to reach out to our support team.</p><p style="font-size: 14px; line-height: 24px; color: #000000;margin:0 0 20px 0;">Thank you for using <span style="font-size: 16px;">SpurtCMS.</span></p><p style="font-size: 16px; line-height: normal; color: #000000;margin:0 0 15px 0;">Best Regards,</p><p style="font-size: 16px; line-height:normal; color: #000000;margin:0 0 10px 0;"><b>SpurtCMS Team</b></p></td></tr>',0,'current-time',1, 0,1,'OTP Generation',1)

INSERT INTO tbl_email_templates(id, template_slug,template_subject,template_description,template_message,module_id,created_on,created_by,is_deleted,is_active,template_name,tenant_id)VALUES(3,'Forgot Password','Reset password link','Forgot Password email template, which sends a confirmation and reset password link to users.', '<tr><td style="padding:0.5rem 1.5rem 0rem;"><h2 style="font-family: "Lexend", sans-serif;font-size: .875rem;font-weight: 400;line-height: 1.125rem;letter-spacing: 0em;text-align: left;color: #525252;margin-bottom: 1.25rem;">Dear <span style="font-family: "Lexend", sans-serif;font-size: 1rem;font-weight: 500;line-height: 20px;letter-spacing: 0em;text-align: left;color: #152027;">{FirstName},</span></h2></td></tr><tr><td style="padding: 0rem 1.5rem;"><p style="font-family: "Lexend", sans-serif;font-size: 14px;font-weight: 400;line-height: 18px;letter-spacing: 0em;text-align: left;color: #525252;margin: 0;">You recently requested to reset your password for spurtCMS admin login.</p><p style="font-family: "Lexend", sans-serif;font-size: 14px;font-weight: 400;line-height: 18px;letter-spacing: 0em;text-align: left;color: #525252;margin-bottom: 1.5rem">Click the link below to reset it.</p></td></tr><tr><td style="padding: 0rem 1.5rem;"><a href="{Passwordurl}" style="color: #2FACD6;text-decoration: underline;">{Passwordurlpath}</a><p class="mb-1" style="font-family: "Lexend", sans-serif;font-size: 14px;font-weight: 400;line-height: 18px;letter-spacing: 0em;text-align: left;color: #525252;">If you did not request a password reset, please ignore this email.</p></td></tr><tr><td style=" padding: 0rem 1.5rem 0rem;"><h2 style="font-family: "Lexend", sans-serif;font-size: .75rem;font-weight: 400;line-height: .9375rem;letter-spacing: 0em;text-align: left;color: #525252;margin-bottom: .5rem;">Best Regards,</h2><p style="font-family: "Lexend", sans-serif;font-size: .875rem;font-weight: 400;line-height: 1.125rem;letter-spacing: 0em;text-align: left;color: #152027;">  Spurt CMS Admin</p></td></tr></tr>' ,0,'current-time',1, 0,1,'Forgot Password',1)

INSERT INTO tbl_email_templates(id, template_slug,template_subject,template_description,template_message,module_id,created_on,created_by,is_deleted,is_active,template_name,tenant_id)VALUES(4,'Createmember','Member registration','Notifies member that they have been created as a member on the platform.','<tr><td><p style="margin-left:0;">&nbsp;</p><h1 style="font-size: 20px; font-weight: bold;line-height: 27px;color:#000000; margin:0 0  12px 0;">Dear <strong>{FirstName}</strong>,</h1></td></tr><tr><td><p style="color:#000000;font-size:14px;">Were glad to have you as a member at spurtCMS. Your membership has been confirmed.</p><p style="color:#000000;font-size:14px;">Log in at</p><p  style="color:#000000;font-size:14px;">{Loginurl}</p><p  style="color:#000000;font-size:14px;margin:0 0 16px;">and indulge in Member related activities and explore our exclusive content.</p></td></tr><tr><td><p style="color:#000000;font-size:16px;line-height:normal;margin:0 0 12px;">Best Regards,</p><p style="color:#000000;font-size:16px;font-weight:500;line-height:24px;margin:0 0 16px;"><strong>Spurt CMS Admin</strong></p></td></tr>' ,5,'current-time',1, 0,1,'Send Member Login credentials',1)

INSERT INTO tbl_email_templates(id, template_slug,template_subject,template_description,template_message,module_id,created_on,created_by,is_deleted,is_active,template_name,tenant_id)VALUES(5,'Logined successfully','User login successfully','User conformation account is logged in successfully','<tr><td><p style="margin-left:0;">&nbsp;</p><h1 style="font-size: 20px; font-weight: bold;line-height: 27px;color:#000000; margin:0 0  12px 0;">Dear <strong>{FirstName}</strong>,</h1></td></tr><tr><td><p style="color:#000000;font-size:14px;">Congratulations! Your SpurtCMS account has been logined successfully.</p><p  style="color:#000000;font-size:14px;margin:0 0 16px;">Start using your Admin Account.</p></td></tr><tr><td><p style="color:#000000;font-size:16px;line-height:normal;margin:0 0 12px;">Best Regards,</p><p style="color:#000000;font-size:16px;font-weight:500;line-height:24px;margin:0 0 16px;"><strong>Spurt CMS Admin</strong></p></td></tr>',0,'current-time',1, 0,1,'Send User Login Email',1)

