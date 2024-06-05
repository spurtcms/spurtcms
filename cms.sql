--Insert Default Values

INSERT INTO tbl_roles(id, name, description, slug, is_active, is_deleted, created_on, created_by) VALUES (1, 'Super Admin', 'Default user type. Has the full administration power', 'super_admin', 1, 0, '2023-07-25 05:50:14', 1);

INSERT INTO tbl_categories(id, category_name, category_slug, created_on, created_by,is_deleted, parent_id, description)	VALUES (1, 'Default Category', 'default_category', '2024-03-04 11:22:03', 1, 0, 0, 'Default_Category'),(2, 'Default1', 'default1', '2024-03-04 11:22:03', 1, 0, 1, 'Default1');

INSERT INTO tbl_channels(id, channel_name, slug_name, field_group_id, is_active, is_deleted, created_on, created_by, channel_description) VALUES (1, 'Default_Channel', 'default_channel', 0, 1, 0, '2024-03-04 10:49:17', '1', 'default description');

INSERT INTO tbl_channel_categories(id, channel_id, category_id, created_at, created_on) VALUES (1, 1, '1,2', 1, '2024-03-04 10:49:17');

INSERT INTO tbl_languages(id,language_name,language_code,json_path,is_status,is_default,created_by,created_on,is_deleted) VALUES (1,'English', 'en', 'locales/en.json', 1, 1,1, '2023-09-11 11:27:44',0)

INSERT INTO tbl_languages(id,language_name,language_code,json_path,is_status,is_default,created_by,created_on,is_deleted) VALUES (2,'Spanish', 'es', 'locales/es.json', 1, 0,1, '2023-09-11 11:27:44',0)

INSERT INTO tbl_languages(id,language_name,language_code,json_path,is_status,is_default,created_by,created_on,is_deleted) VALUES (3,'French', 'fr', 'locales/fr.json', 1, 0,1, '2023-09-11 11:27:44',0)

INSERT INTO tbl_users(id, role_id, first_name, last_name, email, username, password, mobile_no, is_active, created_on, created_by, is_deleted, default_language_id)	VALUES (1, 1, 'Spurt', 'Cms', 'spurtcms@gmail.com', 'Admin', '$2a$14$3kdNuP2Fo/SBopGK9e/9ReBa8uq82vMM5Ko2jnbLpuPb6qxqgR0x2', '9999999999', 1, '2023-11-24 14:56:12',1, 0, 1);

INSERT INTO tbl_user_personalizes(id,user_id,menu_background_color,logo_path,expand_logo_path) values(1,1,'rgb(9, 171, 217)','/public/img/logo.svg','/public/img/logo-bg.svg');

INSERT INTO tbl_member_groups(id,name,slug,description,is_active,is_deleted,created_on,created_by) VALUES (1, 'Default Group', 'default-group', '', 1,0, '2023-11-24 14:56:12', 1);

INSERT INTO tbl_storage_types(id, local, selected_type) VALUES (1, 'storage', 'local');

INSERT INTO tbl_email_configurations(id, selected_type) VALUES (1, 'environment');

INSERT INTO tbl_general_settings(id, logo_path, expand_logo_path, date_format, time_format,language_id,time_zone)	VALUES (1,'/public/img/logo.svg', '/public/img/logo-bg.svg', 'dd mmm yyyy', '12',1,'Asia/Kolkata');

INSERT INTO tbl_member_settings(id, allow_registration,member_login,notification_users)VALUES(1,1,'password','1')

--Channel default fields

INSERT INTO tbl_field_types(id, type_name, type_slug, is_active, is_deleted, created_by, created_on) VALUES (1, 'Label', 'label', 1,  0, 1, '2023-03-14 11:09:12');

INSERT INTO tbl_field_types(id, type_name, type_slug, is_active, is_deleted, created_by, created_on)	VALUES (2, 'Text', 'text', 1,  0, 1, '2023-03-14 11:09:12');

INSERT INTO tbl_field_types(id, type_name, type_slug, is_active, is_deleted, created_by, created_on) VALUES (3, 'Link', 'link', 1,  0, 1, '2023-03-14 11:09:12');
	
INSERT INTO tbl_field_types(id, type_name, type_slug, is_active, is_deleted, created_by, created_on)	VALUES (4, 'Date & Time', 'date&time', 1,  0, 1, '2023-03-14 11:09:12');

INSERT INTO tbl_field_types(id, type_name, type_slug, is_active, is_deleted, created_by, created_on)	VALUES (5, 'Select', 'select', 1,  0, 1, '2023-03-14 11:09:12');

INSERT INTO tbl_field_types(id, type_name, type_slug, is_active, is_deleted, created_by, created_on)	VALUES (6, 'Date', 'date', 1,  0, 1, '2023-03-14 11:09:12');

INSERT INTO tbl_field_types(id, type_name, type_slug, is_active, is_deleted, created_by, created_on)	VALUES (7, 'TextBox', 'textbox', 1,  0, 1, '2023-03-14 11:09:12');

INSERT INTO tbl_field_types(id, type_name, type_slug, is_active, is_deleted, created_by, created_on)	VALUES (8, 'TextArea', 'textarea', 1, 0, 1, '2023-03-14 11:09:12');
	
INSERT INTO tbl_field_types(id, type_name, type_slug, is_active, is_deleted, created_by, created_on)	VALUES (9, 'Radio Button', 'radiobutton', 1, 0, 1, '2023-03-14 11:09:12');
	
INSERT INTO tbl_field_types(id, type_name, type_slug, is_active, is_deleted, created_by, created_on)	VALUES (10, 'CheckBox', 'checkbox', 1, 0, 1, '2023-03-14 11:09:12');
	
INSERT INTO tbl_field_types(id, type_name, type_slug, is_active, is_deleted, created_by, created_on) VALUES (11, 'Text Editor', 'texteditor', 1, 0, 1, '2023-03-14 11:09:12');
	
INSERT INTO tbl_field_types(id, type_name, type_slug, is_active, is_deleted, created_by, created_on)	VALUES (12, 'Section', 'section', 1, 0, 1, '2023-03-14 11:09:12');
	
INSERT INTO tbl_field_types(id, type_name, type_slug, is_active, is_deleted, created_by, created_on)	VALUES (13, 'Section Break', 'sectionbreak', 1, 0, 1, '2023-03-14 11:09:12');

INSERT INTO tbl_field_types(id, type_name, type_slug, is_active, is_deleted, created_by, created_on) VALUES (14, 'Members', 'member', 1,  0, 1, '2023-03-14 11:09:12');

--Default Insert Menu value 

INSERT INTO tbl_modules(id, module_name, is_active, created_by, created_on, default_module, parent_id, assign_permission, icon_path, description, order_index, menu_type,full_access_permission,group_flg) VALUES (1, 'Dashboard', 1, 1, '2023-03-14 11:09:12', 0, 0, 1, '/public/img/dashboard_new.svg', '', 1, 'left',0,0);

INSERT INTO tbl_modules(id, module_name, is_active, created_by, created_on, default_module, parent_id, assign_permission, icon_path, description, order_index, menu_type,full_access_permission,group_flg) VALUES (2, 'Categories', 1, 1, '2023-03-14 11:09:12', 0, 0, 1, '/public/img/category_new.svg', 'Create various categories for content and organize them efficiently by categorizing it with user-defined labels.', 2, 'left',0,0);

INSERT INTO tbl_modules(id, module_name, is_active, created_by, created_on, default_module, parent_id, assign_permission, icon_path, description, order_index, menu_type,full_access_permission,group_flg) VALUES(3, 'Channels', 1, 1, '2023-03-14 11:09:12', 0, 0, 1, '/public/img/cms-on.svg', 'Easily create and manage diverse content types using Spurt CMS Admin.', 3, 'left',0,0);

INSERT INTO tbl_modules(id, module_name, is_active, created_by, created_on, default_module, parent_id, assign_permission, icon_path, description, order_index, menu_type,full_access_permission,group_flg) VALUES(4, 'Media', 1, 1, '2023-03-14 11:09:12', 0, 0, 1, '/public/img/media_1.svg', 'Manage and integrate media files effortlessly to enrich your content.', 4, 'left',0,0);

INSERT INTO tbl_modules(id, module_name, is_active, created_by, created_on, default_module, parent_id, assign_permission, icon_path, description, order_index, menu_type,full_access_permission,group_flg) VALUES(5, 'Members', 1, 1, '2023-03-14 11:09:12', 0, 0, 1, '/public/img/member_new.svg', 'Admin manages member profiles, including adding new members and assigning them to specific groups.', 5, 'left',0,0);

INSERT INTO tbl_modules(id, module_name, is_active, created_by, created_on, default_module, parent_id, assign_permission, icon_path, description, order_index, menu_type,full_access_permission,group_flg) VALUES(6, 'Settings', 1, 1, '2023-03-14 11:09:12', 0, 0, 1, '/public/img/settings.svg', 'Check relevant boxes for providing access control to this role, in this settings.', 6, 'left',0,1);

INSERT INTO tbl_modules(id, module_name, is_active, created_by, created_on, default_module, parent_id, assign_permission, icon_path, description, order_index, menu_type,full_access_permission,group_flg) VALUES(7, 'Channels', 1, 1, '2023-03-14 11:09:12', 0, 3, 0, '/public/img/channels.svg', 'Easily create and manage diverse content types using Spurt CMS Admin.', 8, 'tab',1,0);

INSERT INTO tbl_modules(id, module_name, is_active, created_by, created_on, default_module, parent_id, assign_permission, icon_path, description, order_index, menu_type,full_access_permission,group_flg) VALUES(8, 'Entries', 1, 1, '2023-03-14 11:09:12', 0, 3, 0, '/public/img/entries.svg', 'Seamlessly add and update content within specific channels for targeted delivery and right audience. ', 7, 'tab',1,0);

INSERT INTO tbl_modules(id, module_name, is_active, created_by, created_on, default_module, parent_id, assign_permission, icon_path, description, order_index, menu_type,full_access_permission,group_flg) VALUES(9, 'Categories Group', 1, 1, '2023-03-14 11:09:12', 0, 2, 0, '/public/img/categories-group.svg', 'Create various categories for content and organize them efficiently by categorizing it with user-defined labels.', 9, 'tab',0,0);

INSERT INTO tbl_modules(id, module_name, is_active, created_by, created_on, default_module, parent_id, assign_permission, icon_path, description, order_index, menu_type,full_access_permission,group_flg) VALUES(10, 'Categories', 1, 1, '2023-03-14 11:09:12', 0, 2, 0, '/public/img/categories-group.svg', '', 10, 'tab',0,0);

INSERT INTO tbl_modules(id, module_name, is_active, created_by, created_on, default_module, parent_id, assign_permission, icon_path, description, order_index, menu_type,full_access_permission,group_flg) VALUES(11, 'Member', 1, 1, '2023-03-14 11:09:12', 0, 5, 0, '/public/img/member.svg', 'Add member profiles, map them to a member group and manage the entire list of member profiles. ', 11, 'tab',0,0);

INSERT INTO tbl_modules(id, module_name, is_active, created_by, created_on, default_module, parent_id, assign_permission, icon_path, description, order_index, menu_type,full_access_permission,group_flg) VALUES(12, 'Members Group', 1, 1, '2023-03-14 11:09:12', 0, 5, 0, '/public/img/group.svg', 'Create groups and categorize members into various groups like Elite Members or Favorite Members', 12, 'tab',0,0);

INSERT INTO tbl_modules(id, module_name, is_active, created_by, created_on, default_module, parent_id, assign_permission, icon_path, description, order_index, menu_type,full_access_permission,group_flg) VALUES(13, 'Member Restrict', 1, 1, '2023-03-14 11:09:12', 0, 5, 0, '/public/img/memberAccess.png', '', 13, 'tab',1,0);

INSERT INTO tbl_modules(id, module_name, is_active, created_by, created_on, default_module, parent_id, assign_permission, icon_path, description, order_index, menu_type,full_access_permission,group_flg) VALUES(14, 'My Account', 1, 1, '2023-03-14 11:09:12', 0, 6, 1, '/public/img/profile.svg', 'Here you can change your email id and password.', 14, 'tab',1,0)

INSERT INTO tbl_modules(id, module_name, is_active, created_by, created_on, default_module, parent_id, assign_permission, icon_path, description, order_index, menu_type,full_access_permission,group_flg) VALUES(15, 'Personalize', 1, 1, '2023-03-14 11:09:12', 0, 6, 1, '/public/img/personalize.svg', 'You can change the theme according your preference. ', 15, 'tab',1,0)

INSERT INTO tbl_modules(id, module_name, is_active, created_by, created_on, default_module, parent_id, assign_permission, icon_path, description, order_index, menu_type,full_access_permission,group_flg) VALUES(16, 'Security', 1, 1, '2023-03-14 11:09:12', 0, 6, 1, '/public/img/security.svg', 'To protech your account, you can change your password at regular intervals.', 16, 'tab',1,0)

INSERT INTO tbl_modules(id, module_name, is_active, created_by, created_on, default_module, parent_id, assign_permission, icon_path, description, order_index, menu_type,full_access_permission,group_flg) VALUES(17, 'Roles & Permissions', 1, 1, '2023-03-14 11:09:12', 0, 6, 0, '/public/img/roles-permision.svg', 'Create different admin roles and assign permissions to them based on their responsibilities.', 17, 'tab',1,0)

INSERT INTO tbl_modules(id, module_name, is_active, created_by, created_on, default_module, parent_id, assign_permission, icon_path, description, order_index, menu_type,full_access_permission,group_flg) VALUES(18, 'Team', 1, 1, '2023-03-14 11:09:12', 0, 6, 0, '/public/img/team.svg', 'Create a team of different admin roles.', 18, 'tab',1,0)

INSERT INTO tbl_modules(id, module_name, is_active, created_by, created_on, default_module, parent_id, assign_permission, icon_path, description, order_index, menu_type,full_access_permission,group_flg) VALUES(19, 'Languages', 1, 1, '2023-03-14 11:09:12', 0, 6, 0, '/public/img/language.svg', 'Create and maintain master list of language of preferences for viewing content on the screen in preferred language.', 19, 'tab',1,0)

INSERT INTO tbl_modules(id, module_name, is_active, created_by, created_on, default_module, parent_id, assign_permission, icon_path, description, order_index, menu_type,full_access_permission,group_flg) VALUES(20, 'Email', 1, 1, '2023-03-14 11:09:12', 0, 6, 0, '/public/img/email.svg', 'Create and Maintain Email Templates that are to be sent to users in different scenarios.', 20, 'tab',1,0)

INSERT INTO tbl_modules(id, module_name, is_active, created_by, created_on, default_module, parent_id, assign_permission, icon_path, description, order_index, menu_type,full_access_permission,group_flg) VALUES(21, 'Data', 1, 1, '2023-03-14 11:09:12', 0, 6, 0, '/public/img/settings.svg', '', 21, 'tab',1,0)

INSERT INTO tbl_modules(id, module_name, is_active, created_by, created_on, default_module, parent_id, assign_permission, icon_path, description, order_index, menu_type,full_access_permission,group_flg) VALUES(22, 'Dashboard', 1, 1, '2023-03-14 11:09:12', 0, 1, 1, '/public/img/settings.svg', '', 22, 'tab',1,0);

INSERT INTO tbl_modules(id, module_name, is_active, created_by, created_on, default_module, parent_id, assign_permission, icon_path, description, order_index, menu_type,full_access_permission,group_flg) VALUES(23, 'Media', 1, 1, '2023-03-14 11:09:12', 0, 4, 1, '/public/img/media.svg', '', 23, 'tab',1,0)

INSERT INTO tbl_modules(id, module_name, is_active, created_by, created_on, default_module, parent_id, assign_permission, icon_path, description, order_index, menu_type,full_access_permission,group_flg) VALUES(24, 'Settings', 1, 1, '2023-03-14 11:09:12', 0, 6, 1, '/public/img/settings.svg', '', 24, 'tab',1,0)

-- INSERT INTO tbl_modules(id, module_name, is_active, created_by, created_on, default_module, parent_id, assign_permission, icon_path, description, order_index, menu_type,full_access_permission,group_flg) VALUES(25, 'Settings', 1, 1, '2023-03-14 11:09:12', 0, 2, 1, '/public/img/settings.svg', '', 25, 'tab',1,0)

-- INSERT INTO tbl_modules(id, module_name, is_active, created_by, created_on, default_module, parent_id, assign_permission, icon_path, description, order_index, menu_type,full_access_permission,group_flg) VALUES(26, 'Settings', 1, 1, '2023-03-14 11:09:12', 0, 3, 1, '/public/img/settings.svg', '', 26, 'tab',1,0)

INSERT INTO tbl_modules(id, module_name, is_active, created_by, created_on, default_module, parent_id, assign_permission, icon_path, description, order_index, menu_type,full_access_permission,group_flg) VALUES(27, 'Settings', 1, 1, '2023-03-14 11:09:12', 0, 4, 1, '/public/img/settings.svg', '', 27, 'tab',1,0)

INSERT INTO tbl_modules(id, module_name, is_active, created_by, created_on, default_module, parent_id, assign_permission, icon_path, description, order_index, menu_type,full_access_permission,group_flg) VALUES(28, 'Settings', 1, 1, '2023-03-14 11:09:12', 0, 5, 1, '/public/img/settings.svg', '', 28, 'tab',1,0)

INSERT INTO tbl_modules(id, module_name, is_active, created_by, created_on, default_module, parent_id, assign_permission, icon_path, description, order_index, menu_type,full_access_permission,group_flg) VALUES(29, 'General Settings', 1, 1, '2023-03-14 11:09:12', 0, 6, 0, '/public/img/settings.svg', '', 1, 'tab',1,0)

-- INSERT INTO tbl_modules(id, module_name, is_active, created_by, created_on, default_module, parent_id, assign_permission, icon_path, description, order_index, menu_type,full_access_permission,group_flg) VALUES(30, 'Graphql Api', 1, 1, '2023-03-14 11:09:12', 0, 6, 0, '/public/img/settings.svg', '', 2, 'tab',1,0)

--Default Module Permission Routes

INSERT INTO tbl_module_permissions(id, route_name, display_name, description, module_id, created_by, created_on, full_access_permission, parent_id, assign_permission,order_index, slug_name) VALUES (1, '/dashboard', 'Dashboard', '', 22, 1, '2023-03-14 11:09:12', 1, 0, 1, 1, 'dashboard')

INSERT INTO tbl_module_permissions(id, route_name, display_name, description, module_id, created_by, created_on, full_access_permission, parent_id, assign_permission,order_index, slug_name) VALUES (2, '/categories/', 'View', 'Give view access to the category group', 9, 1, '2023-03-14 11:09:12', 0, 0, 1, 1, 'view')

INSERT INTO tbl_module_permissions(id, route_name, display_name, description, module_id, created_by, created_on, full_access_permission, parent_id, assign_permission,order_index, slug_name) VALUES (3, '/categories/newcategory', 'Create', 'Give create access to the category Group', 9, 1, '2023-03-14 11:09:12', 0, 0, 1, 2, 'create')

INSERT INTO tbl_module_permissions(id, route_name, display_name, description, module_id, created_by, created_on, full_access_permission, parent_id, assign_permission,order_index, slug_name) VALUES (4, '/categories/updatecategory', 'Update', 'Give update access to the category Group', 9, 1, '2023-03-14 11:09:12', 0, 0, 1, 3, 'update')

INSERT INTO tbl_module_permissions(id, route_name, display_name, description, module_id, created_by, created_on, full_access_permission, parent_id, assign_permission,order_index, slug_name) VALUES (5, '/categories/deletecategory', 'Delete', 'Give delete access to the category Group', 9, 1, '2023-03-14 11:09:12', 0, 0, 1, 4, 'delete')

INSERT INTO tbl_module_permissions(id, route_name, display_name, description, module_id, created_by, created_on, full_access_permission, parent_id, assign_permission,order_index, slug_name) VALUES (6, '/categories/addcategory/', 'View', 'Give view access to the categories', 10, 1, '2023-03-14 11:09:12', 0, 0, 1, 1, 'view')

INSERT INTO tbl_module_permissions(id, route_name, display_name, description, module_id, created_by, created_on, full_access_permission, parent_id, assign_permission,order_index, slug_name) VALUES (7, '/categories/addsubcategory', 'Create', 'Give create access to the categories', 10, 1, '2023-03-14 11:09:12', 0, 0, 1, 2, 'create')

INSERT INTO tbl_module_permissions(id, route_name, display_name, description, module_id, created_by, created_on, full_access_permission, parent_id, assign_permission,order_index, slug_name) VALUES (8, '/categories/editsubcategory', 'Update', 'Give update access to the categories', 10, 1, '2023-03-14 11:09:12', 0, 0, 1, 3, 'update')

INSERT INTO tbl_module_permissions(id, route_name, display_name, description, module_id, created_by, created_on, full_access_permission, parent_id, assign_permission,order_index, slug_name) VALUES (9, '/categories/removecategory', 'Delete', 'Give delete access to the categories', 10, 1, '2023-03-14 11:09:12', 0, 0, 1, 4, 'delete')

INSERT INTO tbl_module_permissions(id, route_name, display_name, description, module_id, created_by, created_on, full_access_permission, parent_id, assign_permission,order_index, slug_name) VALUES (10, '/channels/', 'Channel', 'Give full access to the channels', 7, 1, '2023-03-14 11:09:12', 1, 0, 1, 1, 'channels')

INSERT INTO tbl_module_permissions(id, route_name, display_name, description, module_id, created_by, created_on, full_access_permission, parent_id, assign_permission,order_index, slug_name) VALUES (11, '/membersgroup/', 'View', 'Give view access to the group', 12, 1, '2023-03-14 11:09:12', 0, 0,1, 1, 'view')

INSERT INTO tbl_module_permissions(id, route_name, display_name, description, module_id, created_by, created_on, full_access_permission, parent_id, assign_permission,order_index, slug_name) VALUES (12, '/membersgroup/newgroup', 'Create', '', 12, 1, '2023-03-14 11:09:12', 0, 0, 1, 2, 'create')

INSERT INTO tbl_module_permissions(id, route_name, display_name, description, module_id, created_by, created_on, full_access_permission, parent_id, assign_permission,order_index, slug_name) VALUES (13, '/membersgroup/updategroup', 'Update', '', 12, 1, '2023-03-14 11:09:12', 0, 0, 1, 3, 'update')

INSERT INTO tbl_module_permissions(id, route_name, display_name, description, module_id, created_by, created_on, full_access_permission, parent_id, assign_permission,order_index, slug_name) VALUES (14, '/membersgroup/deletegroup', 'Delete', '', 12, 1, '2023-03-14 11:09:12', 0, 0, 1, 4, 'delete')

INSERT INTO tbl_module_permissions(id, route_name, display_name, description, module_id, created_by, created_on, full_access_permission, parent_id, assign_permission,order_index, slug_name) VALUES (15, '/member/', 'View', '', 11, 1, '2023-03-14 11:09:12', 0, 0, 1, 1, 'view')

INSERT INTO tbl_module_permissions(id, route_name, display_name, description, module_id, created_by, created_on, full_access_permission, parent_id, assign_permission,order_index, slug_name) VALUES (16, '/member/newmember', 'Create', '', 11, 1, '2023-03-14 11:09:12', 0, 0, 1, 2, 'create')

INSERT INTO tbl_module_permissions(id, route_name, display_name, description, module_id, created_by, created_on, full_access_permission, parent_id, assign_permission,order_index, slug_name) VALUES (17, '/member/updatemember', 'Update', '', 11, 1, '2023-03-14 11:09:12', 0, 0, 1, 3, 'update')

INSERT INTO tbl_module_permissions(id, route_name, display_name, description, module_id, created_by, created_on, full_access_permission, parent_id, assign_permission,order_index, slug_name) VALUES (18, '/member/deletemember', 'Delete', '', 11, 1, '2023-03-14 11:09:12', 0, 0, 1, 4, 'delete')

INSERT INTO tbl_module_permissions(id, route_name, display_name, description, module_id, created_by, created_on, full_access_permission, parent_id, assign_permission,order_index, slug_name) VALUES (19, '/memberaccess/', 'Member-Restrict', 'Create and manage content consumption access and restrictions to the Members on the website. ', 13, 1, '2023-03-14 11:09:12', 1, 0, 1, 1, 'memberaccess')

INSERT INTO tbl_module_permissions(id, route_name, display_name, description, module_id, created_by, created_on, full_access_permission, parent_id, assign_permission,order_index, slug_name) VALUES (20, '/channel/entrylist', 'Entries', '', 8, 1, '2023-03-14 11:09:12', 1, 0, 0, 0, 'entries')

INSERT INTO tbl_module_permissions(id, route_name, display_name, description, module_id, created_by, created_on, full_access_permission, parent_id, assign_permission,order_index, slug_name) VALUES (21, '/media/', 'View', '', 23, 1, '2023-03-14 11:09:12', 1, 0, 0, 1, 'view')

INSERT INTO tbl_module_permissions(id, route_name, display_name, description, module_id, created_by, created_on, full_access_permission, parent_id, assign_permission,order_index, slug_name) VALUES (22, '/settings/roles/', 'Roles & Permission', 'Create and manage additional roles based on the needs for their tasks and activities in the control panel.', 17, 1, '2023-03-14 11:09:12', 1, 0, 1, 1, 'roles')

INSERT INTO tbl_module_permissions(id, route_name, display_name, description, module_id, created_by, created_on, full_access_permission, parent_id, assign_permission,order_index, slug_name) VALUES (23, '/settings/users/userlist', 'Teams', 'Add users for assigning them roles, followed by assigning permissions for appropriate access to the CMS.', 18, 1, '2023-03-14 11:09:12', 1, 0, 1, 1, 'users')

INSERT INTO tbl_module_permissions(id, route_name, display_name, description, module_id, created_by, created_on, full_access_permission, parent_id, assign_permission,order_index, slug_name) VALUES (24, '/settings/languages/languagelist', 'Languages', 'Add different languages here, for the Members to consume content in the language of preference on the website.', 19, 1, '2023-03-14 11:09:12', 1, 0, 1, 1, 'language')

INSERT INTO tbl_module_permissions(id, route_name, display_name, description, module_id, created_by, created_on, full_access_permission, parent_id, assign_permission,order_index, slug_name) VALUES (25, '/settings/emails', 'Email Templates', 'Add different languages here, for the Members to consume content in the language of preference on the website.', 20, 1, '2023-03-14 11:09:12', 1, 0, 1, 1, 'email')

INSERT INTO tbl_module_permissions(id, route_name, display_name, description, module_id, created_by, created_on, full_access_permission, parent_id, assign_permission,order_index, slug_name) VALUES (26, '/settings/data', 'Data', 'Add different languages here, for the Members to consume content in the language of preference on the website.', 21, 1, '2023-03-14 11:09:12', 1, 0, 1, 1, 'data')

INSERT INTO tbl_module_permissions(id, route_name, display_name, description, module_id, created_by, created_on, full_access_permission, parent_id, assign_permission,order_index, slug_name) VALUES (27, '/categories/settings', 'Settings', '', 25, 1, '2023-03-14 11:09:12', 1, 0, 0, 1, 'settings')

INSERT INTO tbl_module_permissions(id, route_name, display_name, description, module_id, created_by, created_on, full_access_permission, parent_id, assign_permission,order_index, slug_name) VALUES (28, '/channel/settings', 'Settings', '', 26, 1, '2023-03-14 11:09:12', 1, 0, 0, 1, 'settings')

INSERT INTO tbl_module_permissions(id, route_name, display_name, description, module_id, created_by, created_on, full_access_permission, parent_id, assign_permission,order_index, slug_name) VALUES (29, '/media/settings', 'Settings', '', 27, 1, '2023-03-14 11:09:12', 1, 0, 0, 1, 'settings')

INSERT INTO tbl_module_permissions(id, route_name, display_name, description, module_id, created_by, created_on, full_access_permission, parent_id, assign_permission,order_index, slug_name) VALUES (30, '/member/settings', 'Settings', '', 28, 1, '2023-03-14 11:09:12', 1, 0, 0, 1, 'settings')

INSERT INTO tbl_module_permissions(id, route_name, display_name, description, module_id, created_by, created_on, full_access_permission, parent_id, assign_permission,order_index, slug_name) VALUES (31, '/settings/general-settings/', 'General Settings', 'Add different languages here, for the Members to consume content in the language of preference on the website.', 29, 1, '2023-03-14 11:09:12', 1, 0, 1, 1, 'data')

INSERT INTO tbl_module_permissions(id, route_name, display_name, description, module_id, created_by, created_on, full_access_permission, parent_id, assign_permission,order_index, slug_name) VALUES (32, '/settings/graphql/', 'Graphql Api', 'Add different languages here, for the Members to consume content in the language of preference on the website.', 30, 1, '2023-03-14 11:09:12', 1, 0, 1, 1, 'data')

INSERT INTO tbl_module_permissions(id, route_name, display_name, description, module_id, created_by, created_on, full_access_permission, parent_id, assign_permission,order_index, slug_name) VALUES (33, '/channel/entrylist/1', 'Default Channel', '', 8, 1, '2023-03-14 11:09:12', 1, 0, 1, 1, 'Default Channel')

-- Timezone

INSERT INTO tbl_timezones(id,timezone) VALUES (1,'Africa/Cairo'),(2,'Africa/Johannesburg'),(3,'Africa/Lagos'),(4,'Africa/Nairobi'),(5,'America/Argentina/Buenos_Aires'),(6,'America/Chicago'),(7,'America/Denver'),(8,'America/Los_Angeles'),(9,'America/Mexico_City'),(10,'America/New_York'),(11,'America/Sao_Paulo'),(12,'Asia/Bangkok'),(13,'Asia/Dhaka'),(14,'Asia/Dubai'),(15,'Asia/Hong_Kong'),(16,'Asia/Jakarta'),(17,'Asia/Kolkata'),(18,'Asia/Manila'),(19,'Asia/Seoul'),(20,'Asia/Shanghai'),(21,'Asia/Singapore'),(22,'Asia/Tokyo'),(23,'Australia/Melbourne'),(24,'Australia/Sydney'),(25,'Europe/Amsterdam'),(26,'Europe/Berlin'),(27,'Europe/Istanbul'),(28,'Europe/London'),(29,'Europe/Madrid'),(30,'Europe/Moscow'),(31,'Europe/Paris'),(32,'Europe/Rome'),(33,'Pacific/Auckland'),(34,'Pacific/Honolulu')

--Email Default Template Insert

INSERT INTO tbl_email_templates(id, template_name, template_subject, template_message, template_description, module_id, created_on, created_by, is_deleted, is_active)VALUES (1,'createuser', 'Confirm user registration', '<tr><td style="padding: 1.7188rem 1.5rem;border-bottom: .0625rem solid #F5F5F6;background-color: #FDFDFE;"><a href="javascript:void(0)" style="display: inline-block;"><img src="{AdminLogo}" alt="" style="display: block; height:20px"></a></td></tr><tr><td style="padding:2.5rem 1.5rem 0rem;"><h2 style="font-family: \"Lexend\", sans-serif;font-size: .875rem;font-weight: 400;line-height: 1.125rem;letter-spacing: 0em;text-align: left;color: #525252;margin-bottom: 1.25rem;">Dear <span style="font-family: \"Lexend\", sans-serif;font-size: 1rem;font-weight: 500;line-height: 20px;letter-spacing: 0em;text-align: left;color: #152027;">{FirstName},</span></h2><p style="font-family: \"Lexend\", sans-serif;font-size: 14px;font-weight: 400;line-height: 18px;letter-spacing: 0em;text-align: left;color: #525252;margin-bottom: .75rem;">Congratulations! Your SpurtCMS account has been successfully created.</p></td></tr><tr><td style="padding: 0rem 1.5rem;"><p style="font-family: \"Lexend\", sans-serif;font-size: 14px;font-weight: 400;line-height: 18px;letter-spacing: 0em;text-align: left;margin-bottom: 1.5rem;color: #525252;">Here are your login details:</p></td></tr><tr><td style="padding: 0rem 1.5rem;"><p style="font-family: \"Lexend\", sans-serif;font-size: 14px;font-weight: 400;line-height: 18px;letter-spacing: 0em;text-align: left;margin-bottom: .5rem;color: #525252;">Username: <span style="font-family: \"Lexend\", sans-serif;font-size: 16px;font-weight: 500;line-height: 20px;letter-spacing: 0em;text-align: left;color: #152027;">{UserName}</span></p><p style="font-family: \"Lexend\", sans-serif;font-size: 14px;font-weight: 400;line-height: 18px;letter-spacing: 0em;text-align: left;color: #525252;margin-bottom: 1.5rem;">Password: <span style="font-family: \"Lexend\", sans-serif;font-size: 16px;font-weight: 500;line-height: 20px;letter-spacing: 0em;text-align: left;color: #152027;">{Password}</span></p></td></tr><tr><td style="padding: 0rem 1.5rem;"><p style="font-family: \"Lexend\", sans-serif;font-size: 14px;font-weight: 400;line-height: 18px;letter-spacing: 0em;text-align: left;color: #525252;margin: 0;">Please log in at</p> <a href="{Loginurl}" style="color: #2FACD6;text-decoration: underline;">{Loginurl}</a><p class="mb-1" style="font-family: \"Lexend\", sans-serif;font-size: 14px;font-weight: 400;line-height: 18px;letter-spacing: 0em;text-align: left;color: #525252;">and start using your Admin Account.</p></td></tr><tr><td style=" padding: 0rem 1.5rem 2.5rem;"><h2 style="font-family: \"Lexend\", sans-serif;font-size: .75rem;font-weight: 400;line-height: .9375rem;letter-spacing: 0em;text-align: left;color: #525252;margin-bottom: .5rem;">Best Regards,</h2><p style="font-family: \"Lexend\", sans-serif;font-size: .875rem;font-weight: 400;line-height: 1.125rem;letter-spacing: 0em;text-align: left;color: #152027;">Spurt CMS Admin</p></td></tr><tr><td style="padding: 1.5rem 1.5rem 2.5rem;text-align: center;border-top: 1px solid #EDEDED;"><a href="javascript:void(0)" style="font-family: \"Lexend\", sans-serif;font-size: .75rem;font-weight: 300;line-height: .9375rem;letter-spacing: 0em;text-align: center;color: #525252;text-decoration: none;">Privacy policy • Contact us • Read our blog • Join SpurtCMS</a><div style="margin-top: 1rem;"><a href="javascript:void(0)" style="display: inline-block;"><img style="width:24px" src="{TwitterLogo}" alt="" style="margin-right: 1rem;"></a><a href="javascript:void(0)" style="display: inline-block;"><img style="width:24px" src="{FbLogo}" alt="" style="margin-right: 1rem;"></a><a href="javascript:void(0)" style="display: inline-block;"><img style="width:24px" src="{LinkedinLogo}" alt="" style="margin-right: 1rem;"></a></div></td></tr>', 'Change Password: Enable this toggle to activate the Change Password email template, which confirms to users that their password has been successfully changed.', 0, '2023-03-14 11:09:12', 1, 0, 1);

INSERT INTO tbl_email_templates(id, template_name, template_subject, template_message, template_description, module_id, created_on, created_by, is_deleted, is_active)VALUES (2, 'changepassword', 'New password', '<p>Hello,{firstname} your password updated Successfully. your new password:{password}</p>', 'Change Password: Change Password email template, which confirms to users that their password has been successfully changed.', 0, '2023-03-14 11:09:12', 1, 0, 1);

INSERT INTO tbl_email_templates(id, template_name,template_subject,template_description,template_message,module_id,created_on,created_by,is_deleted,is_active)VALUES(3,'OTPGenerate','Confirm OTPs','OTP email template, which sends a One-Time Password (OTP) to users for verification.','<p>Use the following OTP to reset your password. OTP is valid for 5 minutes</p><p>&nbsp;</p><h3>{OTP}</h3>'  ,'5','2023-11-24 14:56:12',1, 0,1)

INSERT INTO tbl_email_templates(id, template_name,template_subject,template_description,template_message,module_id,created_on,created_by,is_deleted,is_active)VALUES(4,'Forgot Password','Reset password link','Forgot Password email template, which sends a confirmation and reset password link to users.', '<tr><td style="padding: 1.7188rem 1.5rem;border-bottom: .0625rem solid #F5F5F6;background-color: #FDFDFE;"><a href="javascript:void(0)" style="display: inline-block;"><img src="{AdminLogo}" alt="" style="display: block; height:20px"></a></td></tr><tr><td style="padding:2.5rem 1.5rem 0rem;"><h2 style="font-family: "Lexend", sans-serif;font-size: .875rem;font-weight: 400;line-height: 1.125rem;letter-spacing: 0em;text-align: left;color: #525252;margin-bottom: 1.25rem;">Dear <span style="font-family: "Lexend", sans-serif;font-size: 1rem;font-weight: 500;line-height: 20px;letter-spacing: 0em;text-align: left;color: #152027;">{FirstName},</span></h2></td></tr><tr><td style="padding: 0rem 1.5rem;"><p style="font-family: "Lexend", sans-serif;font-size: 14px;font-weight: 400;line-height: 18px;letter-spacing: 0em;text-align: left;color: #525252;margin: 0;">You recently requested to reset your password for spurtCMS admin login.</p><p style="font-family: "Lexend", sans-serif;font-size: 14px;font-weight: 400;line-height: 18px;letter-spacing: 0em;text-align: left;color: #525252;margin-bottom: 1.5rem">Click the link below to reset it.</p></td></tr><tr><td style="padding: 0rem 1.5rem;"><a href="{Passwordurl}" style="color: #2FACD6;text-decoration: underline;">{Passwordurlpath}</a><p class="mb-1" style="font-family: "Lexend", sans-serif;font-size: 14px;font-weight: 400;line-height: 18px;letter-spacing: 0em;text-align: left;color: #525252;">If you did not request a password reset, please ignore this email.</p></td></tr><tr><td style=" padding: 0rem 1.5rem 2.5rem;"><h2 style="font-family: "Lexend", sans-serif;font-size: .75rem;font-weight: 400;line-height: .9375rem;letter-spacing: 0em;text-align: left;color: #525252;margin-bottom: .5rem;">Best Regards,</h2><p style="font-family: "Lexend", sans-serif;font-size: .875rem;font-weight: 400;line-height: 1.125rem;letter-spacing: 0em;text-align: left;color: #152027;">  Spurt CMS Admin</p></td></tr><tr><td style="padding: 1.5rem 1.5rem 2.5rem;text-align: center;border-top: 1px solid #EDEDED;"><a href="javascript:void(0)" style="font-family: "Lexend", sans-serif;font-size: .75rem;font-weight: 300;line-height: .9375rem;letter-spacing: 0em;text-align: center;color: #525252;text-decoration: none;">Privacy policy • Contact us • Read our blog • Join SpurtCMS</a><div style="margin-top: 1rem;"><a href="javascript:void(0)" style="display: inline-block;"><img style="width:24px" src="{TwitterLogo}" alt="" style="margin-right: 1rem;"></a> <a href="javascript:void(0)" style="display: inline-block;"><img style="width:24px" src="{FbLogo}" alt=""    style="margin-right: 1rem;"></a><a href="javascript:void(0)" style="display: inline-block;"><img style="width:24px" src="{LinkedinLogo}" alt="" style="margin-right: 1rem;"></a></div></td></tr></tr>' ,0,'2023-11-24 14:56:12',1, 0,1)

INSERT INTO tbl_email_templates(id, template_name,template_subject,template_description,template_message,module_id,created_on,created_by,is_deleted,is_active)VALUES(5,'Createmember','Member registration','Notifies member users that they have been created as a member on the platform.','<figure class="image"><a href="javascript:void(0)"><img style="display:block;height:20px;" src="{AdminLogo}" alt=""></a></figure><h2 style="color:#525252;font-family:"Lexend", sans-serif;font-size:.875rem;font-weight:400;letter-spacing:0em;line-height:1.125rem;margin-bottom:1.25rem;text-align:left;">Dear <span style="color:#152027;font-family:"Lexend", sans-serif;font-size:1rem;"><span style="font-weight:500;letter-spacing:0em;line-height:20px;text-align:left;">{FirstName},</span></span></h2><p style="color:#525252;font-family:"Lexend", sans-serif;font-size:14px;font-weight:400;letter-spacing:0em;line-height:18px;margin-bottom:.75rem;text-align:left;">We’re glad to have you as a member at spurtCMS. Your membership has been confirmed.</p><p style="color:#525252;font-family:"Lexend", sans-serif;font-size:14px;font-weight:400;letter-spacing:0em;line-height:18px;margin-bottom:1.5rem;text-align:left;">Here are your login details:</p><p style="color:#525252;font-family:"Lexend", sans-serif;font-size:14px;font-weight:400;letter-spacing:0em;line-height:18px;margin-bottom:.5rem;text-align:left;">Email Id: <span style="color:#152027;font-family:"Lexend", sans-serif;font-size:16px;"><span style="font-weight:500;letter-spacing:0em;line-height:20px;text-align:left;">{EmailId}</span></span></p><p style="color:#525252;font-family:"Lexend", sans-serif;font-size:14px;font-weight:400;letter-spacing:0em;line-height:18px;margin-bottom:1.5rem;text-align:left;">Password: <span style="color:#152027;font-family:"Lexend", sans-serif;font-size:16px;"><span style="font-weight:500;letter-spacing:0em;line-height:20px;text-align:left;">{Password}</span></span></p><p style="color:#525252;font-family:"Lexend", sans-serif;font-size:14px;font-weight:400;letter-spacing:0em;line-height:18px;margin:0;text-align:left;">Log in at</p><p><a style="color:#2FACD6;" href="{Loginurl}"><u>{Loginurl}</u></a></p><p class="mb-1" style="color:#525252;font-family:"Lexend", sans-serif;font-size:14px;font-weight:400;letter-spacing:0em;line-height:18px;text-align:left;">and indulge in Member related activities and explore our exclusive content.</p><h2 style="color:#525252;font-family:"Lexend", sans-serif;font-size:.75rem;font-weight:400;letter-spacing:0em;line-height:.9375rem;margin-bottom:.5rem;text-align:left;">Best Regards,</h2><p style="color:#152027;font-family:"Lexend", sans-serif;font-size:.875rem;font-weight:400;letter-spacing:0em;line-height:1.125rem;text-align:left;">Spurt CMS Admin</p><p><a style="color:#525252;font-family:"Lexend", sans-serif;font-size:.75rem;font-weight:300;letter-spacing:0em;line-height:.9375rem;text-align:center;text-decoration:none;" href="javascript:void(0)">Privacy&nbsp;policy&nbsp;•&nbsp;Contact&nbsp;us&nbsp;•&nbsp;Read&nbsp;our&nbsp;blog&nbsp;•&nbsp;Join&nbsp;SpurtCMS</a></p><div style="margin-top:1rem;"><a style="display:inline-block;" href="javascript:void(0)"><img class="image_resized" style="width:24px;" src="{TwitterLogo}" alt=""></a> <a style="display:inline-block;" href="javascript:void(0)"><img class="image_resized" style="width:24px;" src="{FbLogo}" alt=""></a> <a style="display:inline-block;" href="javascript:void(0)"><img class="image_resized" style="width:24px;" src="{LinkedinLogo}" alt=""></a></div>' ,5,'2023-11-24 14:56:12',1, 0,1)
