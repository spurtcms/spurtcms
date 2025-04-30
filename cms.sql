--Insert Default Values


INSERT INTO tbl_roles(id, name, description, slug, is_active, is_deleted, created_on, created_by) VALUES (1, 'Super Admin', 'Default user type. Has the full administration power', 'super_admin', 1, 0, '2025-03-14 11:09:12', 1);

INSERT INTO tbl_languages(id,language_name,language_code,image_path,json_path,is_status,is_default,created_by,created_on,is_deleted) VALUES (1,'English', 'en', '/public/img/in.jpeg','locales/en.json', 1, 1,1, '2025-03-14 11:09:12',0)

INSERT INTO tbl_languages(id,language_name,language_code,image_path,json_path,is_status,is_default,created_by,created_on,is_deleted) VALUES (2,'Spanish', 'es','/public/img/sp.jpeg', 'locales/es.json', 1, 0,1, '2025-03-14 11:09:12',0)

INSERT INTO tbl_languages(id,language_name,language_code,image_path,json_path,is_status,is_default,created_by,created_on,is_deleted) VALUES (3,'French', 'fr','/public/img/fr.jpeg', 'locales/fr.json', 1, 0,1, '2025-03-14 11:09:12',0)

INSERT INTO tbl_languages(id,language_name,language_code,image_path,json_path,is_status,is_default,created_by,created_on,is_deleted) VALUES (4,'Russian', 'ru','/public/img/russia.svg', 'locales/ru.json', 1, 0,1, '2025-03-14 11:09:12',0)

INSERT INTO tbl_user_personalizes(id,user_id,menu_background_color,logo_path,expand_logo_path) values(1,1,'rgb(9, 171, 217)','/public/img/logo.svg','/public/img/logo-bg.svg');

-- INSERT INTO tbl_storage_types(id, local, selected_type, aws) VALUES (1, 'storage', 'aws','awsConfig');


INSERT INTO tbl_storage_types(id, local, selected_type) VALUES (1, 'storage', 'aws');


--Channel default fields

INSERT INTO tbl_field_types(id, type_name, type_slug, is_active, is_deleted, created_by, created_on) VALUES (1, 'Label', 'label', 1,  0, 1, '2025-03-14 11:09:12');

INSERT INTO tbl_field_types(id, type_name, type_slug, is_active, is_deleted, created_by, created_on) VALUES (2, 'Text', 'text', 1,  0, 1, '2025-03-14 11:09:12');

INSERT INTO tbl_field_types(id, type_name, type_slug, is_active, is_deleted, created_by, created_on) VALUES (3, 'Link', 'link', 1,  0, 1, '2025-03-14 11:09:12');

INSERT INTO tbl_field_types(id, type_name, type_slug, is_active, is_deleted, created_by, created_on) VALUES (4, 'Date & Time', 'date&time', 1,  0, 1, '2025-03-14 11:09:12');

INSERT INTO tbl_field_types(id, type_name, type_slug, is_active, is_deleted, created_by, created_on) VALUES (5, 'Select', 'select', 1,  0, 1, '2025-03-14 11:09:12');

INSERT INTO tbl_field_types(id, type_name, type_slug, is_active, is_deleted, created_by, created_on) VALUES (6, 'Date', 'date', 1,  0, 1, '2025-03-14 11:09:12');

INSERT INTO tbl_field_types(id, type_name, type_slug, is_active, is_deleted, created_by, created_on) VALUES (7, 'TextBox', 'textbox', 1,  0, 1, '2025-03-14 11:09:12');

INSERT INTO tbl_field_types(id, type_name, type_slug, is_active, is_deleted, created_by, created_on) VALUES (8, 'TextArea', 'textarea', 1, 0, 1, '2025-03-14 11:09:12');

INSERT INTO tbl_field_types(id, type_name, type_slug, is_active, is_deleted, created_by, created_on) VALUES (9, 'Radio Button', 'radiobutton', 1, 0, 1, '2025-03-14 11:09:12');

INSERT INTO tbl_field_types(id, type_name, type_slug, is_active, is_deleted, created_by, created_on) VALUES (10, 'CheckBox', 'checkbox', 1, 0, 1, '2025-03-14 11:09:12');

INSERT INTO tbl_field_types(id, type_name, type_slug, is_active, is_deleted, created_by, created_on) VALUES (11, 'Text Editor', 'texteditor', 1, 0, 1, '2025-03-14 11:09:12');

INSERT INTO tbl_field_types(id, type_name, type_slug, is_active, is_deleted, created_by, created_on) VALUES (12, 'Section', 'section', 1, 0, 1, '2025-03-14 11:09:12');

INSERT INTO tbl_field_types(id, type_name, type_slug, is_active, is_deleted, created_by, created_on) VALUES (13, 'Section Break', 'sectionbreak', 1, 0, 1, '2025-03-14 11:09:12');

INSERT INTO tbl_field_types(id, type_name, type_slug, is_active, is_deleted, created_by, created_on) VALUES (14, 'Members', 'member', 1,  0, 1, '2025-03-14 11:09:12');

INSERT INTO tbl_field_types(id, type_name, type_slug, is_active, is_deleted, created_by, created_on) VALUES (15, 'Media Gallery', 'mediagallrey', 1,  0, 1, '2025-03-14 11:09:12');

INSERT INTO tbl_field_types(id, type_name, type_slug, is_active, is_deleted, created_by, created_on) VALUES (16, 'Video URL', 'videourl', 1,  0, 1, '2025-03-14 11:09:12');


--Default Insert Menu value

INSERT INTO tbl_modules(id, module_name, is_active, created_by, created_on, default_module, parent_id, assign_permission, icon_path, description, order_index, menu_type,full_access_permission,group_flg) VALUES (1, 'Dashboard', 1, 1, '2025-03-14 11:09:12', 0, 0, 1, '/public/img/dashboard_new.svg', '', 1, 'left',0,0);

INSERT INTO tbl_modules(id, module_name, is_active, created_by, created_on, default_module, parent_id, assign_permission, icon_path, description, order_index, menu_type,full_access_permission,group_flg) VALUES (2, 'Categories', 1, 1, '2025-03-14 11:09:12', 0, 0, 1, '/public/img/category_new.svg', 'Create various categories for content and organize them efficiently by categorizing it with user-defined labels.', 2, 'left',0,0);

INSERT INTO tbl_modules(id, module_name, is_active, created_by, created_on, default_module, parent_id, assign_permission, icon_path, description, order_index, menu_type,full_access_permission,group_flg) VALUES(3, 'Channels', 1, 1, '2025-03-14 11:09:12', 0, 0, 1, '/public/img/cms-on.svg', 'Easily create and manage diverse content types using Spurt CMS Admin.', 3, 'left',0,0);

INSERT INTO tbl_modules(id, module_name, is_active, created_by, created_on, default_module, parent_id, assign_permission, icon_path, description, order_index, menu_type,full_access_permission,group_flg) VALUES(4, 'Entries', 1, 1, '2025-03-14 11:09:12', 0, 0, 1, '/public/img/left-entries.svg', 'Seamlessly add and update content within specific channels for targeted delivery and right audience. ', 4, 'tab',0,0);

INSERT INTO tbl_modules(id, module_name, is_active, created_by, created_on, default_module, parent_id, assign_permission, icon_path, description, order_index, menu_type,full_access_permission,group_flg) VALUES(5, 'Users', 1, 1, '2023-03-14 11:09:12', 0, 0, 1, '/public/img/member_new.svg', 'Admin manages user profiles, including adding new users and assigning them to specific groups.', 5, 'left',0,0);

INSERT INTO tbl_modules(id, module_name, is_active, created_by, created_on, default_module, parent_id, assign_permission, icon_path, description, order_index, menu_type,full_access_permission,group_flg) VALUES(6, 'Settings', 1, 1, '2025-03-14 11:09:12', 0, 0, 1, '/public/img/settings.svg', 'Check relevant boxes for providing access control to this role, in this settings.', 6, 'left',0,1);

INSERT INTO tbl_modules(id, module_name, is_active, created_by, created_on, default_module, parent_id, assign_permission, icon_path, description, order_index, menu_type,full_access_permission,group_flg) VALUES(7, 'Channels', 1, 1, '2025-03-14 11:09:12', 0, 3, 0, '/public/img/accord-channels.svg', 'Easily create and manage diverse content types using Spurt CMS Admin.', 7, 'tab',1,0);

INSERT INTO tbl_modules(id, module_name, is_active, created_by, created_on, default_module, parent_id, assign_permission, icon_path, description, order_index, menu_type,full_access_permission,group_flg) VALUES(8, 'Entries', 1, 1, '2025-03-14 11:09:12', 0, 4, 0, '/public/img/entries.svg', 'Seamlessly add and update content within specific channels for targeted delivery and right audience. ', 8, 'tab',1,0);

INSERT INTO tbl_modules(id, module_name, is_active, created_by, created_on, default_module, parent_id, assign_permission, icon_path, description, order_index, menu_type,full_access_permission,group_flg) VALUES(9, 'Categories Group', 1, 1, '2025-03-14 11:09:12', 0, 2, 0, '/public/img/categories-group.svg', 'Create various categories for content and organize them efficiently by categorizing it with user-defined labels.', 9, 'tab',0,0);

INSERT INTO tbl_modules(id, module_name, is_active, created_by, created_on, default_module, parent_id, assign_permission, icon_path, description, order_index, menu_type,full_access_permission,group_flg) VALUES(10, 'Categories', 1, 1, '2025-03-14 11:09:12', 0, 2, 0, '/public/img/categories-group.svg', '', 10, 'tab',0,0);

INSERT INTO tbl_modules(id, module_name, is_active, created_by, created_on, default_module, parent_id, assign_permission, icon_path, description, order_index, menu_type,full_access_permission,group_flg) VALUES(11, 'Users', 1, 1, '2025-03-14 11:09:12', 0, 5, 0, '/public/img/accord-member.svg', 'Add User profiles, map them to a member group and manage the entire list of User profiles. ', 11, 'tab',0,0);

INSERT INTO tbl_modules(id, module_name, is_active, created_by, created_on, default_module, parent_id, assign_permission, icon_path, description, order_index, menu_type,full_access_permission,group_flg) VALUES(12, 'User Group', 1, 1, '2025-03-14 11:09:12', 0, 5, 0, '/public/img/group.svg', 'Create groups and categorize memUsersbers into various groups like Elite Users or Favorite Users', 12, 'tab',0,0);

INSERT INTO tbl_modules(id, module_name, is_active, created_by, created_on, default_module, parent_id, assign_permission, icon_path, description, order_index, menu_type,full_access_permission,group_flg) VALUES(14, 'My Account', 1, 1, '2025-03-14 11:09:12', 0, 6, 1, '/public/img/profile.svg', 'Here you can change your email id and password.', 14, 'tab',1,0)

INSERT INTO tbl_modules(id, module_name, is_active, created_by, created_on, default_module, parent_id, assign_permission, icon_path, description, order_index, menu_type,full_access_permission,group_flg) VALUES(15, 'Personalize', 1, 1, '2025-03-14 11:09:12', 0, 6, 1, '/public/img/personalize.svg', 'You can change the theme according your preference. ', 15, 'tab',1,0)

INSERT INTO tbl_modules(id, module_name, is_active, created_by, created_on, default_module, parent_id, assign_permission, icon_path, description, order_index, menu_type,full_access_permission,group_flg) VALUES(16, 'Roles & Permissions', 1, 1, '2025-03-14 11:09:12', 0, 6, 0, '/public/img/roles-permision.svg', 'Create different admin roles and assign permissions to them based on their responsibilities.', 16, 'tab',1,0)

INSERT INTO tbl_modules(id, module_name, is_active, created_by, created_on, default_module, parent_id, assign_permission, icon_path, description, order_index, menu_type,full_access_permission,group_flg) VALUES(17, 'Team', 1, 1, '2025-03-14 11:09:12', 0, 6, 0, '/public/img/team.svg', 'Create a team of different admin roles.', 17, 'tab',1,0)

INSERT INTO tbl_modules(id, module_name, is_active, created_by, created_on, default_module, parent_id, assign_permission, icon_path, description, order_index, menu_type,full_access_permission,group_flg) VALUES(18, 'Email', 1, 1, '2025-03-14 11:09:12', 0, 6, 0, '/public/img/email.svg', 'Create and Maintain Email Templates that are to be sent to users in different scenarios.', 18, 'tab',1,0)

INSERT INTO tbl_modules(id, module_name, is_active, created_by, created_on, default_module, parent_id, assign_permission, icon_path, description, order_index, menu_type,full_access_permission,group_flg) VALUES(19, 'Dashboard', 1, 1, '2025-03-14 11:09:12', 0, 1, 1, '/public/img/settings.svg', '', 19, 'tab',1,0);

INSERT INTO tbl_modules(id, module_name, is_active, created_by, created_on, default_module, parent_id, assign_permission, icon_path, description, order_index, menu_type,full_access_permission,group_flg) VALUES(20, 'Settings', 1, 1, '2025-03-14 11:09:12', 0, 6, 1, '/public/img/settings.svg', '', 20, 'tab',1,0)

INSERT INTO tbl_modules(id, module_name, is_active, created_by, created_on, default_module, parent_id, assign_permission, icon_path, description, order_index, menu_type,full_access_permission,group_flg) VALUES(21, 'Settings', 1, 1, '2025-03-14 11:09:12', 0, 3, 1, '/public/img/settings.svg', '', 21, 'tab',1,0)

INSERT INTO tbl_modules(id, module_name, is_active, created_by, created_on, default_module, parent_id, assign_permission, icon_path, description, order_index, menu_type,full_access_permission,group_flg) VALUES(22, 'Settings', 1, 1, '2025-03-14 11:09:12', 0, 5, 1, '/public/img/settings.svg', '', 22, 'tab',1,0)

INSERT INTO tbl_modules(id, module_name, is_active, created_by, created_on, default_module, parent_id, assign_permission, icon_path, description, order_index, menu_type,full_access_permission,group_flg) VALUES(23, 'General Settings', 1, 1, '2025-03-14 11:09:12', 0, 6, 0, '/public/img/settings.svg', '', 23, 'tab',1,0)

INSERT INTO tbl_modules(id, module_name, is_active, created_by, created_on, default_module, parent_id, assign_permission, icon_path, description, order_index, menu_type,full_access_permission,group_flg) VALUES(26, 'Graphql Playground', 1, 1, '2025-03-14 11:09:12', 0, 0, 0, '/public/img/graphql-playground.svg', '', 26, 'left',1,1)

INSERT INTO tbl_modules(id, module_name, is_active, created_by, created_on, default_module, parent_id,assign_permission, icon_path, description, order_index, menu_type,full_access_permission,group_flg) VALUES(27, 'Graphql Playground', 1, 1, '2025-03-14 11:09:12', 0, 24, 0, '/public/img/graphql-playground.svg', '', 27, 'tab',1,0)

INSERT INTO tbl_modules(id, module_name, is_active, created_by, created_on, default_module, parent_id, assign_permission, icon_path, description, order_index, menu_type,full_access_permission,group_flg) VALUES(24, 'Next.js Templates', 1, 1, '2025-03-14 11:09:12', 0, 0, 0, '/public/img/templates.svg', '', 24, 'left',1,1)

INSERT INTO tbl_modules(id, module_name, is_active, created_by, created_on, default_module, parent_id, assign_permission, icon_path, description, order_index, menu_type,full_access_permission,group_flg) VALUES(25,  'Next.js Templates', 1, 1, '2025-03-14 11:09:12', 0, 26, 0, '/public/img/templates.svg', '', 25, 'tab',1,0)

INSERT INTO tbl_modules(id, module_name, is_active, created_by, created_on, default_module, parent_id, assign_permission, icon_path, description, order_index, menu_type,full_access_permission,group_flg) VALUES(28, 'Graphql Api', 1, 1, '2025-03-14 11:09:12', 0, 0, 0, '/public/img/role-graph-ql.svg', '', 28, 'left',1,1)

INSERT INTO tbl_modules(id, module_name, is_active, created_by, created_on, default_module, parent_id, assign_permission, icon_path, description, order_index, menu_type,full_access_permission,group_flg) VALUES(29, 'Graphql Api', 1, 1, '2025-03-14 11:09:12', 0, 28, 0, '/public/img/Graph-QL.svg', '', 29, 'tab',1,0)

INSERT INTO tbl_modules(id, module_name, is_active, created_by, created_on, default_module, parent_id, assign_permission, icon_path, description, order_index, menu_type,full_access_permission,group_flg) VALUES(30, 'Languages', 1, 1, '2025-03-14 11:09:12', 0, 0, 0, '/public/img/language.svg', '', 30, 'left',1,1)

INSERT INTO tbl_modules(id, module_name, is_active, created_by, created_on, default_module, parent_id, assign_permission, icon_path, description, order_index, menu_type,full_access_permission,group_flg) VALUES(31, 'Languages', 1, 1, '2025-03-14 11:09:12', 0, 30, 0, '/public/img/language.svg', '', 31, 'tab',1,0)


INSERT INTO tbl_modules(id, module_name, is_active, created_by, created_on, default_module, parent_id, assign_permission, icon_path, description, order_index, menu_type,full_access_permission,group_flg) VALUES(32, 'Forms Builder', 1, 1, '2023-03-14 11:09:12', 0, 0, 0, '/public/img/Group 17556.svg', '', 32, 'left',1,1)

INSERT INTO tbl_modules(id, module_name, is_active, created_by, created_on, default_module, parent_id, assign_permission, icon_path, description, order_index, menu_type,full_access_permission,group_flg) VALUES(33, 'Forms Builder', 1, 1, '2023-03-14 11:09:12', 0, 32, 0, '/public/img/Group 17556.svg', '', 33, 'tab',1,0)


--Default Module Permission Routes

INSERT INTO tbl_module_permissions(id, route_name, display_name, description, module_id, created_by, created_on, full_access_permission, parent_id, assign_permission,order_index, slug_name) VALUES (1, '/dashboard', 'Dashboard', '', 19, 1, '2025-03-14 11:09:12', 1, 0, 1, 1, 'dashboard')

INSERT INTO tbl_module_permissions(id, route_name, display_name, description, module_id, created_by, created_on, full_access_permission, parent_id, assign_permission,order_index, slug_name) VALUES (2, '/categories/', 'View', 'Give view access to the category group', 9, 1, '2025-03-14 11:09:12', 0, 0, 1, 1, 'view')

INSERT INTO tbl_module_permissions(id, route_name, display_name, description, module_id, created_by, created_on, full_access_permission, parent_id, assign_permission,order_index, slug_name) VALUES (3, '/categories/newcategory', 'Create', 'Give create access to the category Group', 9, 1, '2025-03-14 11:09:12', 0, 0, 1, 2, 'create')

INSERT INTO tbl_module_permissions(id, route_name, display_name, description, module_id, created_by, created_on, full_access_permission, parent_id, assign_permission,order_index, slug_name) VALUES (4, '/categories/updatecategory', 'Update', 'Give update access to the category Group', 9, 1, '2025-03-14 11:09:12', 0, 0, 1, 3, 'update')

INSERT INTO tbl_module_permissions(id, route_name, display_name, description, module_id, created_by, created_on, full_access_permission, parent_id, assign_permission,order_index, slug_name) VALUES (5, '/categories/deletecategory', 'Delete', 'Give delete access to the category Group', 9, 1, '2025-03-14 11:09:12', 0, 0, 1, 4, 'delete')

INSERT INTO tbl_module_permissions(id, route_name, display_name, description, module_id, created_by, created_on, full_access_permission, parent_id, assign_permission,order_index, slug_name) VALUES (6, '/categories/addcategory/','View','Give view access to the categories', 10, 1, '2025-03-14 11:09:12', 0, 0, 1, 1, 'view')

INSERT INTO tbl_module_permissions(id, route_name, display_name, description, module_id, created_by, created_on, full_access_permission, parent_id, assign_permission,order_index, slug_name) VALUES (7, '/categories/addsubcategory', 'Create', 'Give create access to the categories', 10, 1, '2025-03-14 11:09:12', 0, 0, 1, 2, 'create')

INSERT INTO tbl_module_permissions(id, route_name, display_name, description, module_id, created_by, created_on, full_access_permission, parent_id, assign_permission,order_index, slug_name) VALUES (8, '/categories/editsubcategory', 'Update', 'Give update access to the categories', 10, 1, '2025-03-14 11:09:12', 0, 0, 1, 3, 'update')

INSERT INTO tbl_module_permissions(id, route_name, display_name, description, module_id, created_by, created_on, full_access_permission, parent_id, assign_permission,order_index, slug_name) VALUES (9, '/categories/removecategory', 'Delete', 'Give delete access to the categories', 10, 1, '2025-03-14 11:09:12', 0, 0, 1, 4, 'delete')

INSERT INTO tbl_module_permissions(id, route_name, display_name, description, module_id, created_by, created_on, full_access_permission, parent_id, assign_permission,order_index, slug_name) VALUES (10, '/channels/', 'Channel', 'Give full access to the channels', 7, 1, '2025-03-14 11:09:12', 1, 0, 1, 1, 'channels')

INSERT INTO tbl_module_permissions(id, route_name, display_name, description, module_id, created_by, created_on, full_access_permission, parent_id, assign_permission,order_index, slug_name) VALUES (11, '/usergroup/', 'View', 'Give view access to the member group', 12, 1, '2025-03-14 11:09:12', 0, 0,1, 1, 'view')

INSERT INTO tbl_module_permissions(id, route_name, display_name, description, module_id, created_by, created_on, full_access_permission, parent_id, assign_permission,order_index, slug_name) VALUES (12, '/usergroup/newgroup', 'Create', 'Give create access to the member group', 12, 1, '2025-03-14 11:09:12', 0, 0, 1, 2, 'create')

INSERT INTO tbl_module_permissions(id, route_name, display_name, description, module_id, created_by, created_on, full_access_permission, parent_id, assign_permission,order_index, slug_name) VALUES (13, '/usergroup/updategroup', 'Update', 'Give update access to the member group', 12, 1, '2025-03-14 11:09:12', 0, 0, 1, 3, 'update')

INSERT INTO tbl_module_permissions(id, route_name, display_name, description, module_id, created_by, created_on, full_access_permission, parent_id, assign_permission,order_index, slug_name) VALUES (14, '/usergroup/deletegroup', 'Delete', 'Give delete access to the member group', 12, 1, '2025-03-14 11:09:12', 0, 0, 1, 4, 'delete')

INSERT INTO tbl_module_permissions(id, route_name, display_name, description, module_id, created_by, created_on, full_access_permission, parent_id, assign_permission,order_index, slug_name) VALUES (15, '/user/', 'View', 'Give view access to the Member', 11, 1, '2025-03-14 11:09:12', 0, 0, 1, 1, 'view')

INSERT INTO tbl_module_permissions(id, route_name, display_name, description, module_id, created_by, created_on, full_access_permission, parent_id, assign_permission,order_index, slug_name) VALUES (16, '/user/newmember', 'Create', 'Give create access to the Member', 11, 1, '2025-03-14 11:09:12', 0, 0, 1, 2, 'create')

INSERT INTO tbl_module_permissions(id, route_name, display_name, description, module_id, created_by, created_on, full_access_permission, parent_id, assign_permission,order_index, slug_name) VALUES (17, '/user/updatemember', 'Update', 'Give update access to the Member', 11, 1, '2025-03-14 11:09:12', 0, 0, 1, 3, 'update')

INSERT INTO tbl_module_permissions(id, route_name, display_name, description, module_id, created_by, created_on, full_access_permission, parent_id, assign_permission,order_index, slug_name) VALUES (18, '/user/deletemember', 'Delete', 'Give delete access to the Member', 11, 1, '2025-03-14 11:09:12', 0, 0, 1, 4, 'delete')

INSERT INTO tbl_module_permissions(id, route_name, display_name, description, module_id, created_by, created_on, full_access_permission, parent_id, assign_permission,order_index, slug_name) VALUES (19, '/memberaccess/', 'Member-Restrict', 'Create and manage content consumption access and restrictions to the Members on the website. ', 13, 1, '2025-03-14 11:09:12', 1, 0, 1, 1, 'memberaccess')

INSERT INTO tbl_module_permissions(id, route_name, display_name, description, module_id, created_by, created_on, full_access_permission, parent_id, assign_permission,order_index, slug_name) VALUES (20, '/entries/entrylist', 'Entries', '', 8, 1, '2025-03-14 11:09:12', 1, 0, 0, 0, 'entries')

INSERT INTO tbl_module_permissions(id, route_name, display_name, description, module_id, created_by, created_on, full_access_permission, parent_id, assign_permission,order_index, slug_name) VALUES (21, '/settings/roles/', 'Roles & Permission', 'Create and manage additional roles based on the needs for their tasks and activities in the control panel.', 16, 1, '2025-03-14 11:09:12', 1, 0, 1, 1, 'roles')

INSERT INTO tbl_module_permissions(id, route_name, display_name, description, module_id, created_by, created_on, full_access_permission, parent_id, assign_permission,order_index, slug_name) VALUES (22, '/settings/users/userlist', 'Teams', 'Add users for assigning them roles, followed by assigning permissions for appropriate access to the CMS.', 17, 1, '2025-03-14 11:09:12', 1, 0, 1, 1, 'users')

INSERT INTO tbl_module_permissions(id, route_name, display_name, description, module_id, created_by, created_on, full_access_permission, parent_id, assign_permission,order_index, slug_name) VALUES (23, '/settings/emails', 'Email Templates', 'Add different languages here, for the Members to consume content in the language of preference on the website.', 18, 1, '2025-03-14 11:09:12', 1, 0, 1, 1, 'email')

INSERT INTO tbl_module_permissions(id, route_name, display_name, description, module_id, created_by, created_on, full_access_permission, parent_id, assign_permission,order_index, slug_name) VALUES (24, '/member/settings', 'Settings', '', 22, 1, '2025-03-14 11:09:12', 1, 0, 0, 1, 'settings')

INSERT INTO tbl_module_permissions(id, route_name, display_name, description, module_id, created_by, created_on, full_access_permission, parent_id, assign_permission,order_index, slug_name) VALUES (25, '/settings/general-settings/', 'General Settings', 'Add different languages here, for the Members to consume content in the language of preference on the website.', 23, 1, '2025-03-14 11:09:12', 1, 0, 1, 1, 'data')

INSERT INTO tbl_module_permissions(id, route_name, display_name, description, module_id, created_by, created_on, full_access_permission, parent_id, assign_permission,order_index, slug_name) VALUES (26, '/channel/entrylist/1', 'Default Channel', '', 8, 1, '2025-03-14 11:09:12', 1, 0, 1, 1, 'Default Channel')

INSERT INTO tbl_module_permissions(id, route_name, display_name, description, module_id, created_by,created_on, full_access_permission, parent_id, assign_permission,order_index, slug_name) VALUES (27, '/graphql', 'Graphql Api', 'Give full access to the graphql api', 29, 1, '2025-03-14 11:09:12', 1, 0, 1, 1, 'graphql')

INSERT INTO tbl_module_permissions(id, route_name, display_name, description, module_id, created_by, created_on, full_access_permission, parent_id, assign_permission,order_index, slug_name) VALUES(28, '/templates', 'Template', 'Give full access to the template', 27, 1, '2025-03-14 11:09:12', 1, 0, 1, 2, 'template')

INSERT INTO tbl_module_permissions(id, route_name, display_name, description, module_id, created_by, created_on, full_access_permission, parent_id, assign_permission,order_index, slug_name) VALUES (29, '/graphql/playground', 'Graphql Playground', 'Give full access to the graphql playground', 25, 1, '2025-03-14 11:09:12', 1, 0, 1, 2, 'graphqlplayground')

INSERT INTO tbl_module_permissions(id, route_name, display_name, description, module_id, created_by, created_on, full_access_permission, parent_id, assign_permission,order_index, slug_name) VALUES (30, '/languages', 'Languages', 'Give full access to the languages', 31, 1, '2025-03-14 11:09:12', 1, 0, 1, 2, 'languages')

INSERT INTO tbl_module_permissions(id, route_name, display_name, description, module_id, created_by, created_on, full_access_permission, parent_id, assign_permission,order_index, slug_name) VALUES (31, '/formsbuilder', 'Forms Builder', 'Give full access to the form builder', 33, 1, '2023-03-14 11:09:12', 1, 0, 1, 2, 'forms-builder')


INSERT INTO tbl_timezones(id,timezone) VALUES (1,'Africa/Cairo'),(2,'Africa/Johannesburg'),(3,'Africa/Lagos'),(4,'Africa/Nairobi'),(5,'America/Argentina/Buenos_Aires'),(6,'America/Chicago'),(7,'America/Denver'),(8,'America/Los_Angeles'),(9,'America/Mexico_City'),(10,'America/New_York'),(11,'America/Sao_Paulo'),(12,'Asia/Bangkok'),(13,'Asia/Dhaka'),(14,'Asia/Dubai'),(15,'Asia/Hong_Kong'),(16,'Asia/Jakarta'),(17,'Asia/Kolkata'),(18,'Asia/Manila'),(19,'Asia/Seoul'),(20,'Asia/Shanghai'),(21,'Asia/Singapore'),(22,'Asia/Tokyo'),(23,'Australia/Melbourne'),(24,'Australia/Sydney'),(25,'Europe/Amsterdam'),(26,'Europe/Berlin'),(27,'Europe/Istanbul'),(28,'Europe/London'),(29,'Europe/Madrid'),(30,'Europe/Moscow'),(31,'Europe/Paris'),(32,'Europe/Rome'),(33,'Pacific/Auckland'),(34,'Pacific/Honolulu')

INSERT INTO tbl_users(id, role_id, first_name, username, password,  is_active, created_on, created_by, is_deleted, default_language_id,tenant_id,s3_folder_name)VALUES (1, 2, 'spurtCMSAdmin',  'spurtcmsAdmin', '$2a$14$3kdNuP2Fo/SBopGK9e/9ReBa8uq82vMM5Ko2jnbLpuPb6qxqgR0x2', 1, '2025-03-14 11:09:12',1, 0, 1,1,'SpurtCMS_1/');

INSERT INTO tbl_roles(id, name, description, slug, is_active, is_deleted, created_on, created_by,tenant_id) VALUES (2, 'Admin', 'Default user type. Has the full administration power', 'admin', 1, 0, '2025-03-14 11:09:12', 1,1);

INSERT INTO tbl_mstr_tenants(id,tenant_id,s3_storage_path,is_deleted)VALUES(1,1,'SpurtCMS_1/',0)

-- Must use separate insert queries for parent and child data insertions to maintain sequence.

-- Some reserved words are used for dynamic values insertion in queries.They are given as below.Use them whenever they needed essentially.

-- 2025-03-14 11:09:12,2(user id),2(tenant id),pid(parent id),chid(channel id),blid(block id),tagid(tag id),mapcat(mapped categories)


INSERT INTO tbl_channels(id,channel_name, slug_name, field_group_id, is_active, is_deleted, created_on, created_by,channel_description,channel_type,tenant_id) VALUES (1,'Default Channel', 'default_channel', 0, 1, 0, '2025-03-14 11:09:12', 1, 'default description','mychannels',1);

INSERT INTO tbl_channel_categories(id,channel_id, category_id, created_at, created_on,tenant_id) VALUES (1,1,'1,2', 1, '2025-03-14 11:09:12',1);

INSERT INTO tbl_member_groups(id,name,slug,description,is_active,is_deleted,created_on,created_by,tenant_id) VALUES (1,'Default Group', 'default-group', 'Default Group', 1,0, '2025-03-14 11:09:12', 1,1);

INSERT INTO tbl_member_settings (id,allow_registration,member_login,notification_users,tenant_id) VALUES (1,1,'password','1',1)

INSERT INTO tbl_email_configurations (id,selected_type,tenant_id) VALUES (1,'environment',1)

INSERT INTO tbl_general_settings (id,date_format,expand_logo_path,language_id,logo_path,storage_type,tenant_id,time_format,time_zone) VALUES (2,'dd mmm yyyy','/public/img/logo-bg.svg',1,'/public/img/logo1.svg','aws',1,'12','Asia/Kolkata')

INSERT INTO tbl_graphql_settings (id,token_name,description,duration,created_by,created_on,is_deleted,token,is_default,tenant_id) VALUES (1,'Default Token','Default Token','Unlimited',1,'2025-03-14 11:09:12',0,'apikey',1,1)

-- below insert queries are for CREATING DEFAULT CATEGORIES, CHANNELS and TEMPLATES 

-- ****************** Knowledge-Base ***********************
--  default knowledge base categories for channel based templates

-- default knowledge base channels for channel based templates

INSERT INTO tbl_module_permissions(id,route_name,display_name,slug_name,module_id,full_access_permission,parent_id,assign_permission,order_index,created_by,created_on,tenant_id) VALUES (31,'/channel/entrylist/2','Knowledge Base','knowledge-base',8,1,0,1,2,1,'2025-03-14 11:09:12',1)




-- ***************************** Blogs *********************************
-- Please DON'T change the flow of the input statements it will definitely affect the output.

--  default Parent blog CATEGORY for channel based templates

-- default Parent blog CHANNEL for channel based templates

INSERT INTO tbl_module_permissions(id,route_name,display_name,slug_name,module_id,full_access_permission,parent_id,assign_permission,order_index,created_by,created_on,tenant_id) VALUES (32,'/channel/entrylist/3','Blog','blog',8,1,0,1,2,1,'2025-03-14 11:09:12',1)



INSERT INTO tbl_email_templates(id, template_slug, template_subject, template_message, template_description, module_id, created_on, created_by, is_deleted, is_active,template_name,tenant_id)VALUES (1,'createuser', 'Confirm user registration', '<tr><td><p style="margin-left:0;">&nbsp;</p><h1 style="font-size: 20px; font-weight: bold;line-height: 27px;color:#000000; margin:0 0  12px 0;">Dear <strong>{FirstName}</strong>,</h1></td></tr><tr><td><p style="color:#000000;font-size:14px;">Congratulations! Your SpurtCMS account has been successfully created.</p><p style="color:#000000;font-size:14px;">Log in at</p><p  style="color:#000000;font-size:14px;">{Loginurl}</p><p  style="color:#000000;font-size:14px;margin:0 0 16px;">and start using your Admin Account.</p></td></tr><tr><td><p style="color:#000000;font-size:16px;line-height:normal;margin:0 0 12px;">Best Regards,</p><p style="color:#000000;font-size:16px;font-weight:500;line-height:24px;margin:0 0 16px;"><strong>Spurt CMS Admin</strong></p></td></tr>', 'Change Password: Enable this toggle to activate the Change Password email template, which confirms to users that their password has been successfully changed.', 0, '2025-03-14 11:09:12', 1, 0, 1,'Create Admin User',1);

INSERT INTO tbl_email_templates(id, template_slug, template_subject, template_message, template_description, module_id, created_on, created_by, is_deleted, is_active,template_name,tenant_id)VALUES (2, 'changepassword', 'New password', '<p>Hello,{firstname} your password updated Successfully. your new password:{password}</p>', 'Change Password: Change Password email template, which confirms to users that their password has been successfully changed.', 0, '2025-03-14 11:09:12', 1, 0, 1,'Change Password',1);

-- INSERT INTO tbl_email_templates(id, template_slug,template_subject,template_description,template_message,module_id,created_on,created_by,is_deleted,is_active,template_name,tenant_id)VALUES(3,'OTPGenerate','Confirm OTPs','OTP email template, which sends a One-Time Password (OTP) to users for verification.','<tr><td><h1 style="font-size: 20px; font-weight: bold;line-height: 27px;color:#000000; margin:0 0  16px 0;">Dear {FirstName}</h1></td></tr><tr><td><p style="font-size: 14px; line-height: 24px; color: #000000;margin:0;">As requsted, here is your One-Time Password (OTP) to log in to SpurtCMS:</p><h3 style="font-size: 20px; line-height:1; font-weight: bold; color: #000000; margin:20px 0;">OTP:{OTP}</h3><p style="font-size: 14px; line-height: 24px; color: #000000;margin:0 0 20px 0;">Please use this OTP within <span style="font-size: 16px; font-weight:bold;">{Expirytime}</span> to complete the login process. If you do not log in within this time frame, you may request a new OTP.</p><p style="font-size: 14px; line-height: 24px; color: #000000;margin:0 0 20px 0;">If you encounter any issue or need assistance, feel free to reach out to our support team.</p><p style="font-size: 14px; line-height: 24px; color: #000000;margin:0 0 20px 0;">Thank you for using <span style="font-size: 16px;">SpurtCMS.</span></p><p style="font-size: 16px; line-height: normal; color: #000000;margin:0 0 15px 0;">Best Regards,</p><p style="font-size: 16px; line-height:normal; color: #000000;margin:0 0 10px 0;"><b>SpurtCMS Team</b></p></td></tr>',0,'2025-03-14 11:09:12',1, 0,1,'OTP Generation',1)

INSERT INTO tbl_email_templates(id, template_slug,template_subject,template_description,template_message,module_id,created_on,created_by,is_deleted,is_active,template_name,tenant_id)VALUES(3,'Forgot Password','Reset password link','Forgot Password email template, which sends a confirmation and reset password link to users.', '<tr><td style="padding:0.5rem 1.5rem 0rem;"><h2 style="font-family: "Lexend", sans-serif;font-size: .875rem;font-weight: 400;line-height: 1.125rem;letter-spacing: 0em;text-align: left;color: #525252;margin-bottom: 1.25rem;">Dear <span style="font-family: "Lexend", sans-serif;font-size: 1rem;font-weight: 500;line-height: 20px;letter-spacing: 0em;text-align: left;color: #152027;">{FirstName},</span></h2></td></tr><tr><td style="padding: 0rem 1.5rem;"><p style="font-family: "Lexend", sans-serif;font-size: 14px;font-weight: 400;line-height: 18px;letter-spacing: 0em;text-align: left;color: #525252;margin: 0;">You recently requested to reset your password for spurtCMS admin login.</p><p style="font-family: "Lexend", sans-serif;font-size: 14px;font-weight: 400;line-height: 18px;letter-spacing: 0em;text-align: left;color: #525252;margin-bottom: 1.5rem">Click the link below to reset it.</p></td></tr><tr><td style="padding: 0rem 1.5rem;"><a href="{Passwordurl}" style="color: #2FACD6;text-decoration: underline;">{Passwordurlpath}</a><p class="mb-1" style="font-family: "Lexend", sans-serif;font-size: 14px;font-weight: 400;line-height: 18px;letter-spacing: 0em;text-align: left;color: #525252;">If you did not request a password reset, please ignore this email.</p></td></tr><tr><td style=" padding: 0rem 1.5rem 0rem;"><h2 style="font-family: "Lexend", sans-serif;font-size: .75rem;font-weight: 400;line-height: .9375rem;letter-spacing: 0em;text-align: left;color: #525252;margin-bottom: .5rem;">Best Regards,</h2><p style="font-family: "Lexend", sans-serif;font-size: .875rem;font-weight: 400;line-height: 1.125rem;letter-spacing: 0em;text-align: left;color: #152027;">  Spurt CMS Admin</p></td></tr></tr>' ,0,'2025-03-14 11:09:12',1, 0,1,'Forgot Password',1)

INSERT INTO tbl_email_templates(id, template_slug,template_subject,template_description,template_message,module_id,created_on,created_by,is_deleted,is_active,template_name,tenant_id)VALUES(4,'Createmember','Member registration','Notifies member that they have been created as a member on the platform.','<tr><td><p style="margin-left:0;">&nbsp;</p><h1 style="font-size: 20px; font-weight: bold;line-height: 27px;color:#000000; margin:0 0  12px 0;">Dear <strong>{FirstName}</strong>,</h1></td></tr><tr><td><p style="color:#000000;font-size:14px;">Were glad to have you as a member at spurtCMS. Your membership has been confirmed.</p><p style="color:#000000;font-size:14px;">Log in at</p><p  style="color:#000000;font-size:14px;">{Loginurl}</p><p  style="color:#000000;font-size:14px;margin:0 0 16px;">and indulge in Member related activities and explore our exclusive content.</p></td></tr><tr><td><p style="color:#000000;font-size:16px;line-height:normal;margin:0 0 12px;">Best Regards,</p><p style="color:#000000;font-size:16px;font-weight:500;line-height:24px;margin:0 0 16px;"><strong>Spurt CMS Admin</strong></p></td></tr>' ,5,'2025-03-14 11:09:12
',1, 0,1,'Send Member Login credentials',1)

INSERT INTO tbl_email_templates(id, template_slug,template_subject,template_description,template_message,module_id,created_on,created_by,is_deleted,is_active,template_name,tenant_id)VALUES(5,'Logined successfully','User login successfully','User conformation account is logged in successfully','<tr><td><p style="margin-left:0;">&nbsp;</p><h1 style="font-size: 20px; font-weight: bold;line-height: 27px;color:#000000; margin:0 0  12px 0;">Dear <strong>{FirstName}</strong>,</h1></td></tr><tr><td><p style="color:#000000;font-size:14px;">Congratulations! Your SpurtCMS account has been logined successfully.</p><p  style="color:#000000;font-size:14px;margin:0 0 16px;">Start using your Admin Account.</p></td></tr><tr><td><p style="color:#000000;font-size:16px;line-height:normal;margin:0 0 12px;">Best Regards,</p><p style="color:#000000;font-size:16px;font-weight:500;line-height:24px;margin:0 0 16px;"><strong>Spurt CMS Admin</strong></p></td></tr>',0,'2025-03-14 11:09:12',1, 0,1,'Send User Login Email',1)


-- Blogs

INSERT INTO tbl_template_modules(id,template_module_name,template_module_slug,description,is_active,created_by,created_on,slug_name) VALUES (1,'Blog', 'blog', 'Blog Templates', 1, 1, 'current_time','blog')

INSERT INTO tbl_templates(id, template_name, template_slug, template_module_id, image_name, image_path, deploy_link, github_link, created_on, created_by, is_active, is_deleted,slug_name,preview) VALUES(1,'Contentflow','contentflow_next_js_theme', 1,'blogstarter.png','/public/img/blogstarter.png','https://vercel.com/new/clone?repository-url=https://github.com/spurtcms/nextjs-starter-theme&demo-title=nextjs-starter-theme&env=NEXT_PUBLIC_SPURTCMS_NEXTJS_STARTER_APIKEY','https://github.com/spurtcms/nextjs-starter-theme','current_time',1,1,0,'blog','https://nextjs-starter-theme.vercel.app/')

INSERT INTO tbl_templates(id,template_name, template_slug, template_module_id, image_name, image_path, deploy_link, github_link, created_on, created_by, is_active, is_deleted,slug_name,preview) VALUES(2,'Writer`s hub theme','nextjs_writer`s_hub_theme', 1,'blog1.png','/public/img/blog1.png','https://vercel.com/new/clone?repository-url=https://github.com/spurtcms/nextjs-blog-theme&demo-title=nextjs-writer`s-hub-theme&env=NEXT_PUBLIC_SPURTCMS_NEXTJS_STARTER_APIKEY','https://github.com/spurtcms/nextjs-blog-theme','current_time',1,1,0,'blog','https://nextjs-blog-theme-liart.vercel.app/')

INSERT INTO tbl_templates(id,template_name, template_slug, template_module_id, image_name, image_path, deploy_link, github_link, created_on, created_by, is_active, is_deleted,slug_name,preview) VALUES(3,'Storyze','next-js-blog2-theme', 1,'blog2.png','/public/img/blog2.png','https://vercel.com/new/clone?repository-url=https://github.com/spurtcms/nextjs-blog2-theme&demo-title=nextjs-blog2-theme&env=NEXT_PUBLIC_SPURTCMS_NEXTJS_STARTER_APIKEY','https://github.com/spurtcms/nextjs-blog2-theme','current_time',1,1,0,'blog','https://storyze-nextjs-theme.vercel.app/')

INSERT INTO tbl_templates( id,template_name, template_slug, template_module_id, image_name, image_path, deploy_link, github_link, created_on, created_by, is_active, is_deleted,slug_name) VALUES(4,'Content chronicle','content_chronicle_nextJS_theme', 1,'blog4.png','/public/img/blog4.png','https://vercel.com/new/clone?repository-url=https://github.com/spurtcms/nextjs-blog4-theme&demo-title=nextjs-blog4-theme&env=NEXT_PUBLIC_SPURTCMS_NEXTJS_STARTER_APIKEY','https://github.com/spurtcms/nextjs-blog4-theme','current_time',1,1,0,'blog')

INSERT INTO tbl_templates( id,template_name, template_slug, template_module_id, image_name, image_path, deploy_link, github_link, created_on, created_by, is_active, is_deleted,slug_name) VALUES(5,'InsightSphere','insightsphere', 1,'V1-blog1.jpeg','/public/img/V1-blog1.jpeg','https://vercel.com/new/clone?repository-url=https://github.com/spurtcms/v1-blog1-theme&demo-title=v1-blog1-theme&env=NEXT_PUBLIC_SPURTCMS_NEXTJS_STARTER_APIKEY','https://github.com/spurtcms/v1-blog1-theme','current_time',1,1,0,'blog')

INSERT INTO tbl_templates( id,template_name, template_slug, template_module_id, image_name, image_path, deploy_link, github_link, created_on, created_by, is_active, is_deleted,slug_name,preview) VALUES(6,'ContentVerse','contentverse', 1,'V2-blog2.jpeg','/public/img/V2-blog2.jpeg','https://vercel.com/new/clone?repository-url=https://github.com/spurtcms/v1-blog2-theme&demo-title=v1-blog2-theme&env=NEXT_PUBLIC_SPURTCMS_NEXTJS_STARTER_APIKEY','https://github.com/spurtcms/v1-blog2-theme','current_time',1,1,0,'blog','https://content-verse-five.vercel.app/')


-- knowledgeBase

INSERT INTO tbl_template_modules(id,template_module_name,template_module_slug,description,is_active,created_by,created_on,slug_name) VALUES ( 2,'Help Center', 'help-center', 'Help Center Templates', 1, 1, 'current_time','knowledge-base')

INSERT INTO tbl_templates(id,template_name, template_slug, template_module_id, image_name, image_path, deploy_link, github_link, created_on, created_by, is_active, is_deleted,slug_name,preview) VALUES(7,'Docuno','docuno-nextjs-theme', 2,'knowledgebase.png','/public/img/knowledgebase.png','https://vercel.com/new/clone?repository-url=https://github.com/spurtcms/knowledge-base&demo-title=knowledge-base&env=NEXT_PUBLIC_SPURTCMS_NEXTJS_STARTER_APIKEY','https://github.com/spurtcms/knowledge-base','current_time',1,1,0,'knowledge-base','https://knowledge-base-lyart-one.vercel.app/')

INSERT INTO tbl_templates(id,template_name, template_slug, template_module_id, image_name, image_path, deploy_link, github_link, created_on, created_by, is_active, is_deleted,slug_name,preview) VALUES(8,'Support Sphere','support-sphere', 2,'knowledgebasenew.jpg','/public/img/knowledgebasenew.jpg','https://vercel.com/new/clone?repository-url=https://github.com/spurtcms/help-center&demo-title=knowledge-base&env=NEXT_PUBLIC_SPURTCMS_NEXTJS_STARTER_APIKEY','https://github.com/spurtcms/help-center','current_time',1,1,0,'knowledge-base','https://support-sphere-gilt.vercel.app/')

-- News

INSERT INTO tbl_template_modules(id,template_module_name,template_module_slug,description,is_active,created_by,created_on,slug_name) VALUES ( 3,'News', 'news', 'News Templates', 1, 1, 'current_time','news')

INSERT INTO tbl_templates(id,template_name, template_slug, template_module_id, image_name, image_path, deploy_link, github_link, created_on, created_by, is_active, is_deleted,slug_name) VALUES(9,'Newsprism','newsprism-nextjs-theme', 3,'CoverimageNewsTemplate.jpeg','/public/img/CoverimageNewsTemplate.jpeg','https://vercel.com/new/clone?repository-url=https://github.com/spurtcms/newsprism-nextjs-theme&demo-title=newsprism-nextjs-theme&env=NEXT_PUBLIC_SPURTCMS_NEXTJS_STARTER_APIKEY','https://github.com/spurtcms/Newsprism-nextjs-theme','current_time',1,1,0,'news')

-- Jobs

INSERT INTO tbl_template_modules(id,template_module_name,template_module_slug,description,is_active,created_by,created_on,slug_name) VALUES (4,'Jobs', 'jobs', 'Jobs Templates', 1, 1, 'current_time','jobs')

INSERT INTO tbl_templates(id,template_name, template_slug, template_module_id, image_name, image_path, deploy_link, github_link, created_on, created_by, is_active, is_deleted,slug_name,preview) VALUES(10,'JobiFylo','jobifylo-nextjs-theme', 4,'jobstemplate.png','/public/img/jobstemplate.png','https://vercel.com/new/clone?repository-url=https://github.com/spurtcms/jobifylo-nextjs-theme&demo-title=jobifylo-nextjs-theme&env=NEXT_PUBLIC_SPURTCMS_NEXTJS_STARTER_APIKEY','https://github.com/spurtcms/nextjs-jobs-theme','current_time',1,1,0,'jobs','https://jobi-fylo-nextjs-theme.vercel.app/')





INSERT INTO tbl_categories(id,category_name, category_slug, created_on, created_by,is_deleted, parent_id, description, tenant_id)	VALUES (1,'Default Category', 'default_category', 'current_time', 1, 0, 0, 'Default_Category', 1);

INSERT INTO tbl_categories(id,category_name, category_slug, created_on, created_by,is_deleted, parent_id, description, tenant_id)	VALUES (2,'Default1', 'default1', 'current_time', 1, 0, 1, 'Default1',1);




--  default Parent blog CATEGORY for channel based templates
INSERT INTO tbl_categories(id,category_name, category_slug, created_on, created_by, is_deleted, parent_id, description, tenant_id) VALUES (3,'Blog', 'blog', 'current_time', 1, 0,0,'Parent Category for blogs', 1)

INSERT INTO tbl_categories(id,category_name, category_slug, created_on, created_by,is_deleted, parent_id, description, tenant_id)	VALUES (4,'Technology', 'technology', 'current_time', 1, 0, 3, 'Technology',1);

INSERT INTO tbl_categories(id,category_name, category_slug, created_on, created_by,is_deleted, parent_id, description, tenant_id)	VALUES (5,'Lifestyle', 'lifestyle', 'current_time', 1, 0, 3, 'Lifestyle',1);

INSERT INTO tbl_categories(id,category_name, category_slug, created_on, created_by,is_deleted, parent_id, description, tenant_id)	VALUES (6,'Travel', 'travel', 'current_time', 1, 0, 3, 'Travel',1);

INSERT INTO tbl_categories(id,category_name, category_slug, created_on, created_by,is_deleted, parent_id, description, tenant_id)	VALUES (7,'Business', 'business', 'current_time', 1, 0, 3, 'Business',1);

INSERT INTO tbl_categories(id,category_name, category_slug, created_on, created_by,is_deleted, parent_id, description, tenant_id)	VALUES (8,'Education', 'education', 'current_time', 1, 0, 3, 'Education',1);


-- Ecommerce


INSERT INTO tbl_categories(id,category_name, category_slug, created_on, created_by,is_deleted, parent_id, description, tenant_id)	VALUES (9,'Ecommerce', 'ecommerce', 'current_time', 1, 0, 0, 'Ecommerce', 1);

INSERT INTO tbl_categories(id,category_name, category_slug, created_on, created_by,is_deleted, parent_id, description, tenant_id)	VALUES (10,'Fashion & Clothing', 'fashion_clothing', 'current_time', 1, 0, 9, 'Fashion & Clothing',1);

    INSERT INTO tbl_categories(id,category_name, category_slug, created_on, created_by,is_deleted, parent_id, description, tenant_id)	VALUES (11,'Men`s Clothing', 'mens_clothing', 'current_time', 1, 0, 10, 'Men`s Clothing',1);
    INSERT INTO tbl_categories(id,category_name, category_slug, created_on, created_by,is_deleted, parent_id, description, tenant_id)	VALUES (12,'Women`s Clothing', 'womens_clothing', 'current_time', 1, 0, 10, 'Women`s Clothing',1);
    INSERT INTO tbl_categories(id,category_name, category_slug, created_on, created_by,is_deleted, parent_id, description, tenant_id)	VALUES (13,'Kid`s Clothing', 'sids_clothing', 'current_time', 1, 0, 10, 'Kid`s Clothing',1);
    INSERT INTO tbl_categories(id,category_name, category_slug, created_on, created_by,is_deleted, parent_id, description, tenant_id)	VALUES (14,'Accessories', 'accessories', 'current_time', 1, 0, 10, 'Accessories',1);
    INSERT INTO tbl_categories(id,category_name, category_slug, created_on, created_by,is_deleted, parent_id, description, tenant_id)	VALUES (15,'Footwear', 'footwear', 'current_time', 1, 0, 10, 'Footwear',1);


INSERT INTO tbl_categories(id,category_name, category_slug, created_on, created_by,is_deleted, parent_id, description, tenant_id)	VALUES (16,'Electronics', 'electronics', 'current_time', 1, 0, 9, 'Electronics',1);

    INSERT INTO tbl_categories(id,category_name, category_slug, created_on, created_by,is_deleted, parent_id, description, tenant_id)	VALUES (17,'Mobile Phones', 'mobile_phones', 'current_time', 1, 0, 16, 'Mobile Phones',1);
    INSERT INTO tbl_categories(id,category_name, category_slug, created_on, created_by,is_deleted, parent_id, description, tenant_id)	VALUES (18,'Laptops & Computers', 'laptops_computers', 'current_time', 1, 0, 16, 'Laptops & Computers',1);
    INSERT INTO tbl_categories(id,category_name, category_slug, created_on, created_by,is_deleted, parent_id, description, tenant_id)	VALUES (19,'Cameras', 'cameras', 'current_time', 1, 0, 16, 'Cameras',1);
    INSERT INTO tbl_categories(id,category_name, category_slug, created_on, created_by,is_deleted, parent_id, description, tenant_id)	VALUES (20,'Home Appliances', 'home_appliances', 'current_time', 1, 0, 16, 'Home Appliances',1);


INSERT INTO tbl_categories(id,category_name, category_slug, created_on, created_by,is_deleted, parent_id, description, tenant_id)	VALUES (21,'Home & Furniture', 'home_furniture', 'current_time', 1, 0, 9, 'Home & Furniture',1);

    INSERT INTO tbl_categories(id,category_name, category_slug, created_on, created_by,is_deleted, parent_id, description, tenant_id)	VALUES (22,'Living Room', 'living_room', 'current_time', 1, 0, 21, 'Living Room',1);
    INSERT INTO tbl_categories(id,category_name, category_slug, created_on, created_by,is_deleted, parent_id, description, tenant_id)	VALUES (23,'Bedroom', 'bedroom', 'current_time', 1, 0, 21, 'Bedroom',1);
    INSERT INTO tbl_categories(id,category_name, category_slug, created_on, created_by,is_deleted, parent_id, description, tenant_id)	VALUES (24,'Kitchen', 'kitchen', 'current_time', 1, 0, 21, 'Kitchen',1);
    INSERT INTO tbl_categories(id,category_name, category_slug, created_on, created_by,is_deleted, parent_id, description, tenant_id)	VALUES (25,'Decor', 'decor', 'current_time', 1, 0, 21, 'Decor',1);

INSERT INTO tbl_categories(id,category_name, category_slug, created_on, created_by,is_deleted, parent_id, description, tenant_id)	VALUES (26,'Beauty & Personal Care', 'beauty_personal_care', 'current_time', 1, 0, 9, 'Beauty & Personal Care',1);

    INSERT INTO tbl_categories(id,category_name, category_slug, created_on, created_by,is_deleted, parent_id, description, tenant_id)	VALUES (27,'Makeup', 'makeup', 'current_time', 1, 0, 26, 'Makeup',1);
    INSERT INTO tbl_categories(id,category_name, category_slug, created_on, created_by,is_deleted, parent_id, description, tenant_id)	VALUES (28,'Skincare', 'skincare', 'current_time', 1, 0, 26, 'Skincare',1);
    INSERT INTO tbl_categories(id,category_name, category_slug, created_on, created_by,is_deleted, parent_id, description, tenant_id)	VALUES (29,'Haircare', 'haircare', 'current_time', 1, 0, 26, 'Haircare',1);
    INSERT INTO tbl_categories(id,category_name, category_slug, created_on, created_by,is_deleted, parent_id, description, tenant_id)	VALUES (30,'Fragrances', 'fragrances', 'current_time', 1, 0, 26, 'Fragrances',1);



INSERT INTO tbl_categories(id,category_name, category_slug, created_on, created_by,is_deleted, parent_id, description, tenant_id)	VALUES (31,'Sports & Outdoors', 'sports_outdoors', 'current_time', 1, 0, 9, 'Sports & Outdoors',1);

    INSERT INTO tbl_categories(id,category_name, category_slug, created_on, created_by,is_deleted, parent_id, description, tenant_id)	VALUES (32,'Fitness Equipment', 'fitness_equipment', 'current_time', 1, 0, 31, 'Fitness Equipment',1);
    INSERT INTO tbl_categories(id,category_name, category_slug, created_on, created_by,is_deleted, parent_id, description, tenant_id)	VALUES (33,'Sports Gear', 'sports_gear', 'current_time', 1, 0, 31, 'Sports Gear',1);
    INSERT INTO tbl_categories(id,category_name, category_slug, created_on, created_by,is_deleted, parent_id, description, tenant_id)	VALUES (34,'Activewear', 'activewear', 'current_time', 1, 0, 31, 'Activewear',1);
    INSERT INTO tbl_categories(id,category_name, category_slug, created_on, created_by,is_deleted, parent_id, description, tenant_id)	VALUES (35,'Outdoor', 'outdoor', 'current_time', 1, 0, 31, 'Outdoor',1);


-- jobs

INSERT INTO tbl_categories(id,category_name, category_slug, created_on, created_by,is_deleted, parent_id, description, tenant_id)	VALUES (36,'Jobs', 'jobs', 'current_time', 1, 0, 0, 'Jobs', 1);

INSERT INTO tbl_categories(id,category_name, category_slug, created_on, created_by,is_deleted, parent_id, description, tenant_id)	VALUES (37,'Healthcare', 'healthcare', 'current_time', 1, 0, 36, 'Healthcare',1);
INSERT INTO tbl_categories(id,category_name, category_slug, created_on, created_by,is_deleted, parent_id, description, tenant_id)	VALUES (38,'Qualification', 'qualification', 'current_time', 1, 0, 36, 'Education',1);
INSERT INTO tbl_categories(id,category_name, category_slug, created_on, created_by,is_deleted, parent_id, description, tenant_id)	VALUES (39,'Finance', 'finance', 'current_time', 1, 0, 36, 'Finance',1);
INSERT INTO tbl_categories(id,category_name, category_slug, created_on, created_by,is_deleted, parent_id, description, tenant_id)	VALUES (40,'Marketing & Sales', 'marketing_sales', 'current_time', 1, 0, 36, 'Marketing & Sales',1);
INSERT INTO tbl_categories(id,category_name, category_slug, created_on, created_by,is_deleted, parent_id, description, tenant_id)	VALUES (41,'Engineering', 'engineering', 'current_time', 1, 0, 36, 'Engineering',1);

-- news

INSERT INTO tbl_categories(id,category_name, category_slug, created_on, created_by,is_deleted, parent_id, description, tenant_id)	VALUES (42,'News', 'news', 'current_time', 1, 0, 0, 'News', 1);

INSERT INTO tbl_categories(id,category_name, category_slug, created_on, created_by,is_deleted, parent_id, description, tenant_id)	VALUES (43,'Politics', 'politics', 'current_time', 1, 0, 42, 'Politics',1);

    INSERT INTO tbl_categories(id,category_name, category_slug, created_on, created_by,is_deleted, parent_id, description, tenant_id)	VALUES (44,'National', 'national', 'current_time', 1, 0, 43, 'National',1);
    INSERT INTO tbl_categories(id,category_name, category_slug, created_on, created_by,is_deleted, parent_id, description, tenant_id)	VALUES (45,'International', 'international', 'current_time', 1, 0, 43, 'International',1);


INSERT INTO tbl_categories(id,category_name, category_slug, created_on, created_by,is_deleted, parent_id, description, tenant_id)	VALUES (46,'Finance', 'finance', 'current_time', 1, 0, 42, 'Finance',1);

    INSERT INTO tbl_categories(id,category_name, category_slug, created_on, created_by,is_deleted, parent_id, description, tenant_id)	VALUES (47,'Markets', 'markets', 'current_time', 1, 0, 46, 'Markets',1);
    INSERT INTO tbl_categories(id,category_name, category_slug, created_on, created_by,is_deleted, parent_id, description, tenant_id)	VALUES (48,'Economy', 'economy', 'current_time', 1, 0, 46, 'Economy',1);
    INSERT INTO tbl_categories(id,category_name, category_slug, created_on, created_by,is_deleted, parent_id, description, tenant_id)	VALUES (49,'Corporate', 'corporate', 'current_time', 1, 0, 46, 'Corporate',1);


INSERT INTO tbl_categories(id,category_name, category_slug, created_on, created_by,is_deleted, parent_id, description, tenant_id)	VALUES (50,'Tech Insights', 'tech_insights', 'current_time', 1, 0, 42, 'Tech Insights',1);

    INSERT INTO tbl_categories(id,category_name, category_slug, created_on, created_by,is_deleted, parent_id, description, tenant_id)	VALUES (51,'Gadgets', 'gadgets', 'current_time', 1, 0, 50, 'Gadgets',1);
    INSERT INTO tbl_categories(id,category_name, category_slug, created_on, created_by,is_deleted, parent_id, description, tenant_id)	VALUES (52,'Innovation', 'innovation', 'current_time', 1, 0, 50, 'Innovation',1);
    INSERT INTO tbl_categories(id,category_name, category_slug, created_on, created_by,is_deleted, parent_id, description, tenant_id)	VALUES (53,'Apps & Software', 'apps_software', 'current_time', 1, 0, 50, 'Apps & Software',1);


INSERT INTO tbl_categories(id,category_name, category_slug, created_on, created_by,is_deleted, parent_id, description, tenant_id)	VALUES (54,'Health', 'health', 'current_time', 1, 0, 42, 'Health',1);

    INSERT INTO tbl_categories(id,category_name, category_slug, created_on, created_by,is_deleted, parent_id, description, tenant_id)	VALUES (55,'Physical Health', 'physical_health', 'current_time', 1, 0, 54, 'Physical Health',1);
    INSERT INTO tbl_categories(id,category_name, category_slug, created_on, created_by,is_deleted, parent_id, description, tenant_id)	VALUES (56,'Mental Health', 'mental_health', 'current_time', 1, 0, 54, 'Mental Health',1);
    INSERT INTO tbl_categories(id,category_name, category_slug, created_on, created_by,is_deleted, parent_id, description, tenant_id)	VALUES (57,'Healthcare Industry', 'healthcare_industry', 'current_time', 1, 0, 54, 'Healthcare Industry',1);


INSERT INTO tbl_categories(id,category_name, category_slug, created_on, created_by,is_deleted, parent_id, description, tenant_id)	VALUES (58,'Sports', 'sports', 'current_time', 1, 0, 42, 'Sports',1);

    INSERT INTO tbl_categories(id,category_name, category_slug, created_on, created_by,is_deleted, parent_id, description, tenant_id)	VALUES (59,'Cricket', 'cricket', 'current_time', 1, 0, 58, 'Cricket',1);
    INSERT INTO tbl_categories(id,category_name, category_slug, created_on, created_by,is_deleted, parent_id, description, tenant_id)	VALUES (60,'Football', 'football', 'current_time', 1, 0, 58, 'Football',1);
    INSERT INTO tbl_categories(id,category_name, category_slug, created_on, created_by,is_deleted, parent_id, description, tenant_id)	VALUES (61,'Other Sports', 'other_sports', 'current_time', 1, 0, 58, 'Other Sports',1);


INSERT INTO tbl_categories(id,category_name, category_slug, created_on, created_by,is_deleted, parent_id, description, tenant_id)	VALUES (62,'World News', 'worldnews', 'current_time', 1, 0, 42, 'World News',1);

    INSERT INTO tbl_categories(id,category_name, category_slug, created_on, created_by,is_deleted, parent_id, description, tenant_id)	VALUES (63,'Asia', 'asia', 'current_time', 1, 0, 62, 'Asia',1);
    INSERT INTO tbl_categories(id,category_name, category_slug, created_on, created_by,is_deleted, parent_id, description, tenant_id)	VALUES (64,'Middle east', 'Middle East', 'current_time', 1, 0, 62, 'Middle East',1);
    INSERT INTO tbl_categories(id,category_name, category_slug, created_on, created_by,is_deleted, parent_id, description, tenant_id)	VALUES (65,'North america', 'North America', 'current_time', 1, 0, 62, 'North America',1);
    INSERT INTO tbl_categories(id,category_name, category_slug, created_on, created_by,is_deleted, parent_id, description, tenant_id)	VALUES (66,'Europe', 'europe', 'current_time', 1, 0, 62, 'Europe',1);



-- ****************** Knowledge-Base ***********************
--  default knowledge base categories for channel based templates
INSERT INTO tbl_categories(id,category_name, category_slug, created_on, created_by, is_deleted, parent_id, description, tenant_id) VALUES (67,'Knowledge Base', 'knowledge-base', 'current_time', 1, 0,0,'Parent Category for knowledge base', 1)

INSERT INTO tbl_categories(id,category_name, category_slug, created_on, created_by,is_deleted, parent_id, description, tenant_id)	VALUES (68,'Getting Started', 'gettingstarted', 'current_time', 1, 0, 67, 'Getting Started',1);

INSERT INTO tbl_categories(id,category_name, category_slug, created_on, created_by,is_deleted, parent_id, description, tenant_id)	VALUES (69,'Product Documentation', 'productdocumentation', 'current_time', 1, 0, 67, 'Product Documentation',1);

INSERT INTO tbl_categories(id,category_name, category_slug, created_on, created_by,is_deleted, parent_id, description, tenant_id)	VALUES (70,'Troubleshooting', 'troubleshooting', 'current_time', 1, 0, 67, 'Troubleshooting',1);

INSERT INTO tbl_categories(id,category_name, category_slug, created_on, created_by,is_deleted, parent_id, description, tenant_id)	VALUES (71,'Billing & Payments', 'billingpayments', 'current_time', 1, 0, 67, 'Billing & Payments',1);

INSERT INTO tbl_categories(id,category_name, category_slug, created_on, created_by,is_deleted, parent_id, description, tenant_id)	VALUES (72,'Integrations', 'Integrations', 'current_time', 1, 0, 67, 'Integrations',1);

INSERT INTO tbl_categories(id,category_name, category_slug, created_on, created_by,is_deleted, parent_id, description, tenant_id)	VALUES (73,'Release Notes', 'release_notes', 'current_time', 1, 0, 67, 'Release Notes',1);

INSERT INTO tbl_forms(id,form_title,form_slug,form_data,status,is_active,created_by,created_on,is_deleted,uuid,form_image_path,form_description,channel_name,tenant_id,modified_on,channel_id,form_preview_imagepath)VALUES (1,'Connect with Us','connect-with-us','{"data":[{"uniqID":"98c2d620-41fa-446b-bcf4-5853dabeeba9","type":"Title","value":"Connect with Us"},{"type":"ShortAnswer","uniqID":"649fd04f-79ad-4150-a8ce-00032798f53f","question":"Name","placeholder":"Enter Your Name","value":"","required":false,"min":"","max":"","textareaHeight":"4rem"},{"type":"ShortAnswer","uniqID":"9bc4d442-17bb-40f2-a5b2-ef764b8314ee","question":"Email Address","placeholder":"Enter Your Email Id","value":"","required":false,"min":"","max":"","textareaHeight":"4rem"},{"type":"Dropdown","uniqID":"3825d8ee-b526-4f4b-ae95-19ef6e433e67","question":"Reason for Contact","options":[{"value":"General Enquiry","id":1},{"value":"Feedback","id":2},{"value":"Collaboration","id":3},{"value":"Technical Issue","id":4}],"value":"","required":false,"textareaHeight":"4rem"},{"type":"ShortAnswer","uniqID":"39f285c2-f05a-4bbd-93db-a7aa7fc4bbdc","question":"Subject","placeholder":"Enter Your Subject","value":"","required":false,"min":"","max":"","textareaHeight":"4rem"},{"type":"LongAnswer","uniqID":"0d078eb4-49ea-4cb0-bfd0-ace1e75142e0","question":"Message","placeholder":"Enter Your Message","value":"","required":false,"min":"","max":"","textareaHeight":"4rem"}],"submit":"Submit","date":"Wed Jan 22 2025 17:36:07 GMT+0530 (India Standard Time)"}',1,1,1,'current_time',0,'12adcd88c6ef','/public/img/CAT-2.svg','Contact us with your questions or feedback, and we’ll respond promptly! Let me know if you need further adjustments!','Default Channel',1,'current_time',1,'cta/connectus.png')
