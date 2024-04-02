--Insert Default Values

INSERT INTO public.tbl_roles(id, name, description, slug, is_active, is_deleted, created_on, created_by) VALUES (1, 'Super Admin', 'Default user type. Has the full administration power', 'super_admin', 1, 0, '2023-07-25 05:50:14', 1);

INSERT INTO public.tbl_categories(id, category_name, category_slug, created_on, created_by,is_deleted, parent_id, description)	VALUES (1, 'Default Category', 'default_category', '2024-03-04 11:22:03', 1, 0, 0, 'Default_Category'),(2, 'Default1', 'default1', '2024-03-04 11:22:03', 1, 0, 1, 'Default1');

INSERT INTO public.tbl_channels(id, channel_name, slug_name, field_group_id, is_active, is_deleted, created_on, created_by, channel_description) VALUES (1, 'Default_Channel', 'default_channel', 0, 1, 0, '2024-03-04 10:49:17', '1', 'default description');

INSERT INTO public.tbl_channel_categories(id, channel_id, category_id, created_at, created_on) VALUES (1, 1, '1,2', 1, '2024-03-04 10:49:17');

INSERT INTO PUBLIC.tbl_languages(ID,LANGUAGE_NAME,LANGUAGE_CODE,JSON_PATH,IS_STATUS,IS_DEFAULT,	CREATED_BY,CREATED_ON,IS_DELETED) VALUES (1,'English', 'en', 'locales/en.json', 1, 1,1, '2023-09-11 11:27:44',0)

INSERT INTO public.tbl_users(id, role_id, first_name, last_name, email, username, password, mobile_no, is_active, created_on, created_by, is_deleted, default_language_id)	VALUES (1, 1, 'Spurt', 'Cms', 'spurtcms@gmail.com', 'Admin', '$2a$14$3kdNuP2Fo/SBopGK9e/9ReBa8uq82vMM5Ko2jnbLpuPb6qxqgR0x2', '9999999999', 1, '2023-11-24 14:56:12',1, 0, 1);

INSERT INTO PUBLIC.TBL_MEMBER_GROUPS(ID,NAME,SLUG,DESCRIPTION,IS_ACTIVE,IS_DELETED,	CREATED_ON,CREATED_BY) VALUES (1, 'Default Group', 'default-group', '', 1,0, '2023-11-24 14:56:12', 1);

--Channel default fields

INSERT INTO public.tbl_field_types(id, type_name, type_slug, is_active, is_deleted, created_by, created_on) VALUES (1, 'Label', 'label', 1,  0, 1, '2023-03-14 11:09:12');

INSERT INTO public.tbl_field_types(id, type_name, type_slug, is_active, is_deleted, created_by, created_on)	VALUES (2, 'Text', 'text', 1,  0, 1, '2023-03-14 11:09:12');

INSERT INTO public.tbl_field_types(id, type_name, type_slug, is_active, is_deleted, created_by, created_on) VALUES (3, 'Link', 'link', 1,  0, 1, '2023-03-14 11:09:12');
	
INSERT INTO public.tbl_field_types(id, type_name, type_slug, is_active, is_deleted, created_by, created_on)	VALUES (4, 'Date & Time', 'date&time', 1,  0, 1, '2023-03-14 11:09:12');

INSERT INTO public.tbl_field_types(id, type_name, type_slug, is_active, is_deleted, created_by, created_on)	VALUES (5, 'Select', 'select', 1,  0, 1, '2023-03-14 11:09:12');

INSERT INTO public.tbl_field_types(id, type_name, type_slug, is_active, is_deleted, created_by, created_on)	VALUES (6, 'Date', 'date', 1,  0, 1, '2023-03-14 11:09:12');

INSERT INTO public.tbl_field_types(id, type_name, type_slug, is_active, is_deleted, created_by, created_on)	VALUES (7, 'TextBox', 'textbox', 1,  0, 1, '2023-03-14 11:09:12');

INSERT INTO public.tbl_field_types(id, type_name, type_slug, is_active, is_deleted, created_by, created_on)	VALUES (8, 'TextArea', 'textarea', 1, 0, 1, '2023-03-14 11:09:12');
	
INSERT INTO public.tbl_field_types(id, type_name, type_slug, is_active, is_deleted, created_by, created_on)	VALUES (9, 'Radio Button', 'radiobutton', 1, 0, 1, '2023-03-14 11:09:12');
	
INSERT INTO public.tbl_field_types(id, type_name, type_slug, is_active, is_deleted, created_by, created_on)	VALUES (10, 'CheckBox', 'checkbox', 1, 0, 1, '2023-03-14 11:09:12');
	
INSERT INTO public.tbl_field_types(id, type_name, type_slug, is_active, is_deleted, created_by, created_on) VALUES (11, 'Text Editor', 'texteditor', 1, 0, 1, '2023-03-14 11:09:12');
	
INSERT INTO public.tbl_field_types(id, type_name, type_slug, is_active, is_deleted, created_by, created_on)	VALUES (12, 'Section', 'section', 1, 0, 1, '2023-03-14 11:09:12');
	
INSERT INTO public.tbl_field_types(id, type_name, type_slug, is_active, is_deleted, created_by, created_on)	VALUES (13, 'Section Break', 'sectionbreak', 1, 0, 1, '2023-03-14 11:09:12');

INSERT INTO public.tbl_field_types(id, type_name, type_slug, is_active, is_deleted, created_by, created_on) VALUES (14, 'Members', 'member', 1,  0, 1, '2023-03-14 11:09:12');

--Default Insert Menu value 

INSERT INTO public.tbl_modules(id, module_name, is_active, created_by, created_on, default_module, parent_id, assign_permission, icon_path, description, order_index, menu_type) VALUES (1, 'Dashboard', 1, 1, '2023-03-14 11:09:12', 0, 0, 0, '/public/img/dashboard-on.svg', '', 1, 'left');

INSERT INTO public.tbl_modules(id, module_name, is_active, created_by, created_on, default_module, parent_id, assign_permission, icon_path, description, order_index, menu_type) VALUES (2, 'Categories', 1, 1, '2023-03-14 11:09:12', 0, 0, 0, '/public/img/category group.svg', 'Create various categories for content and organize them efficiently by categorizing it with user-defined labels.', 2, 'left');

INSERT INTO public.tbl_modules(id, module_name, is_active, created_by, created_on, default_module, parent_id, assign_permission, icon_path, description, order_index, menu_type) VALUES(3, 'Channels', 1, 1, '2023-03-14 11:09:12', 0, 0, 0, '/public/img/cms-on.svg', 'Easily create and manage diverse content types using Spurt CMS Admin.', 3, 'left');

INSERT INTO public.tbl_modules(id, module_name, is_active, created_by, created_on, default_module, parent_id, assign_permission, icon_path, description, order_index, menu_type) VALUES(4, 'Media', 1, 1, '2023-03-14 11:09:12', 0, 0, 0, '/public/img/media_1.svg', 'Manage and integrate media files effortlessly to enrich your content.', 4, 'left');

INSERT INTO public.tbl_modules(id, module_name, is_active, created_by, created_on, default_module, parent_id, assign_permission, icon_path, description, order_index, menu_type) VALUES(5, 'Members', 1, 1, '2023-03-14 11:09:12', 0, 0, 0, '/public/img/members-on.svg', 'Admin manages member profiles, including adding new members and assigning them to specific groups.', 5, 'left');

INSERT INTO public.tbl_modules(id, module_name, is_active, created_by, created_on, default_module, parent_id, assign_permission, icon_path, description, order_index, menu_type) VALUES(6, 'Settings', 1, 1, '2023-03-14 11:09:12', 0, 0, 0, '/public/img/settings.svg', 'Check relevant boxes for providing access control to this role, in this settings.', 6, 'left');

INSERT INTO public.tbl_modules(id, module_name, is_active, created_by, created_on, default_module, parent_id, assign_permission, icon_path, description, order_index, menu_type) VALUES(7, 'Channels', 1, 1, '2023-03-14 11:09:12', 0, 3, 0, '/public/img/cms-on.svg', 'Easily create and manage diverse content types using Spurt CMS Admin.', 7, 'tab');

INSERT INTO public.tbl_modules(id, module_name, is_active, created_by, created_on, default_module, parent_id, assign_permission, icon_path, description, order_index, menu_type) VALUES(8, 'Entries', 1, 1, '2023-03-14 11:09:12', 0, 3, 0, '/public/img/entries.svg', 'Seamlessly add and update content within specific channels for targeted delivery and right audience. ', 8, 'tab');

INSERT INTO public.tbl_modules(id, module_name, is_active, created_by, created_on, default_module, parent_id, assign_permission, icon_path, description, order_index, menu_type) VALUES(9, 'Categories Group', 1, 1, '2023-03-14 11:09:12', 0, 2, 0, '/public/img/category group.svg', 'Create various categories for content and organize them efficiently by categorizing it with user-defined labels.', 9, 'tab');

INSERT INTO public.tbl_modules(id, module_name, is_active, created_by, created_on, default_module, parent_id, assign_permission, icon_path, description, order_index, menu_type) VALUES(10, 'Categories', 1, 1, '2023-03-14 11:09:12', 0, 2, 0, '/public/img/category group.svg', '', 10, 'tab');

INSERT INTO public.tbl_modules(id, module_name, is_active, created_by, created_on, default_module, parent_id, assign_permission, icon_path, description, order_index, menu_type) VALUES(11, 'Member', 1, 1, '2023-03-14 11:09:12', 0, 5, 0, '/public/img/member.svg', 'Add member profiles, map them to a member group and manage the entire list of member profiles. ', 11, 'tab');

INSERT INTO public.tbl_modules(id, module_name, is_active, created_by, created_on, default_module, parent_id, assign_permission, icon_path, description, order_index, menu_type) VALUES(12, 'Members Group', 1, 1, '2023-03-14 11:09:12', 0, 5, 0, '/public/img/group.svg', 'Create groups and categorize members into various groups like Elite Members or Favorite Members', 12, 'tab');

INSERT INTO public.tbl_modules(id, module_name, is_active, created_by, created_on, default_module, parent_id, assign_permission, icon_path, description, order_index, menu_type) VALUES(13, 'Member Restrict', 1, 1, '2023-03-14 11:09:12', 0, 5, 0, '/public/img/memberAccess.png', '', 13, 'tab');

INSERT INTO public.tbl_modules(id, module_name, is_active, created_by, created_on, default_module, parent_id, assign_permission, icon_path, description, order_index, menu_type) VALUES(14, 'My Account', 1, 1, '2023-03-14 11:09:12', 0, 6, 1, '/public/img/profile.svg', 'Here you can change your email id and password.', 14, 'tab')

INSERT INTO public.tbl_modules(id, module_name, is_active, created_by, created_on, default_module, parent_id, assign_permission, icon_path, description, order_index, menu_type) VALUES(15, 'Personalize', 1, 1, '2023-03-14 11:09:12', 0, 6, 1, '/public/img/personalize.svg', 'You can change the theme according your preference. ', 15, 'tab')

INSERT INTO public.tbl_modules(id, module_name, is_active, created_by, created_on, default_module, parent_id, assign_permission, icon_path, description, order_index, menu_type) VALUES(16, 'Security', 1, 1, '2023-03-14 11:09:12', 0, 6, 1, '/public/img/security.svg', 'To protech your account, you can change your password at regular intervals.', 16, 'tab')

INSERT INTO public.tbl_modules(id, module_name, is_active, created_by, created_on, default_module, parent_id, assign_permission, icon_path, description, order_index, menu_type) VALUES(17, 'Roles & Permissions', 1, 1, '2023-03-14 11:09:12', 0, 6, 0, '/public/img/roles-permision.svg', 'Create different admin roles and assign permissions to them based on their responsibilities.', 17, 'tab')

INSERT INTO public.tbl_modules(id, module_name, is_active, created_by, created_on, default_module, parent_id, assign_permission, icon_path, description, order_index, menu_type) VALUES(18, 'Team', 1, 1, '2023-03-14 11:09:12', 0, 6, 0, '/public/img/team.svg', 'Create a team of different admin roles.', 18, 'tab')

INSERT INTO public.tbl_modules(id, module_name, is_active, created_by, created_on, default_module, parent_id, assign_permission, icon_path, description, order_index, menu_type) VALUES(19, 'Languages', 1, 1, '2023-03-14 11:09:12', 0, 6, 0, '/public/img/language.svg', 'Create and maintain master list of language of preferences for viewing content on the screen in preferred language.', 19, 'tab')

INSERT INTO public.tbl_modules(id, module_name, is_active, created_by, created_on, default_module, parent_id, assign_permission, icon_path, description, order_index, menu_type) VALUES(20, 'Email Templates', 1, 1, '2023-03-14 11:09:12', 0, 6, 0, '/public/img/email.svg', 'Create and Maintain Email Templates that are to be sent to users in different scenarios.', 20, 'tab')

INSERT INTO public.tbl_modules(id, module_name, is_active, created_by, created_on, default_module, parent_id, assign_permission, icon_path, description, order_index, menu_type) VALUES(21, 'Data', 1, 1, '2023-03-14 11:09:12', 0, 6, 0, '', '', 21, 'tab')

INSERT INTO public.tbl_modules(id, module_name, is_active, created_by, created_on, default_module, parent_id, assign_permission, icon_path, description, order_index, menu_type) VALUES(22, 'Dashboard', 1, 1, '2023-03-14 11:09:12', 0, 1, 0, '', '', 22, 'tab');

INSERT INTO public.tbl_modules(id, module_name, is_active, created_by, created_on, default_module, parent_id, assign_permission, icon_path, description, order_index, menu_type) VALUES(23, 'Media', 1, 1, '2023-03-14 11:09:12', 0, 4, 0, '', '', 23, 'tab')

--Default Module Permission Routes

INSERT INTO public.tbl_module_permissions(id, route_name, display_name, description, module_id, created_by, created_on, full_access_permission, parent_id, assign_permission,order_index, slug_name) VALUES (1, '/dashboard', 'Dashboard', '', 22, 1, '2023-03-14 11:09:12', 1, 0, 1, 1, 'dashboard')

INSERT INTO public.tbl_module_permissions(id, route_name, display_name, description, module_id, created_by, created_on, full_access_permission, parent_id, assign_permission,order_index, slug_name) VALUES (2, '/categories/', 'View', '', 9, 1, '2023-03-14 11:09:12', 1, 0, 1, 1, 'view')

INSERT INTO public.tbl_module_permissions(id, route_name, display_name, description, module_id, created_by, created_on, full_access_permission, parent_id, assign_permission,order_index, slug_name) VALUES (3, '/categories/newcategory', 'Create', '', 9, 1, '2023-03-14 11:09:12', 1, 0, 1, 2, 'create')

INSERT INTO public.tbl_module_permissions(id, route_name, display_name, description, module_id, created_by, created_on, full_access_permission, parent_id, assign_permission,order_index, slug_name) VALUES (4, '/categories/updatecategory', 'Update', '', 9, 1, '2023-03-14 11:09:12', 1, 0, 1, 3, 'update')

INSERT INTO public.tbl_module_permissions(id, route_name, display_name, description, module_id, created_by, created_on, full_access_permission, parent_id, assign_permission,order_index, slug_name) VALUES (5, '/categories/deletecategory', 'Delete', '', 9, 1, '2023-03-14 11:09:12', 1, 0, 1, 4, 'delete')

INSERT INTO public.tbl_module_permissions(id, route_name, display_name, description, module_id, created_by, created_on, full_access_permission, parent_id, assign_permission,order_index, slug_name) VALUES (6, '/categories/addcategory', 'View', '', 10, 1, '2023-03-14 11:09:12', 1, 0, 1, 1, 'view')

INSERT INTO public.tbl_module_permissions(id, route_name, display_name, description, module_id, created_by, created_on, full_access_permission, parent_id, assign_permission,order_index, slug_name) VALUES (7, '/categories/addsubcategory', 'Create', '', 10, 1, '2023-03-14 11:09:12', 1, 0, 1, 2, 'create')

INSERT INTO public.tbl_module_permissions(id, route_name, display_name, description, module_id, created_by, created_on, full_access_permission, parent_id, assign_permission,order_index, slug_name) VALUES (8, '/categories/editsubcategory', 'Update', '', 10, 1, '2023-03-14 11:09:12', 1, 0, 1, 3, 'update')

INSERT INTO public.tbl_module_permissions(id, route_name, display_name, description, module_id, created_by, created_on, full_access_permission, parent_id, assign_permission,order_index, slug_name) VALUES (9, '/categories/removecategory', 'Delete', '', 10, 1, '2023-03-14 11:09:12', 1, 0, 1, 4, 'delete')

INSERT INTO public.tbl_module_permissions(id, route_name, display_name, description, module_id, created_by, created_on, full_access_permission, parent_id, assign_permission,order_index, slug_name) VALUES (10, '/channels/', 'Channel', '', 7, 1, '2023-03-14 11:09:12', 1, 0, 1, 1, 'channels')

INSERT INTO public.tbl_module_permissions(id, route_name, display_name, description, module_id, created_by, created_on, full_access_permission, parent_id, assign_permission,order_index, slug_name) VALUES (11, '/membersgroup/', 'View', '', 12, 1, '2023-03-14 11:09:12', 1, 0, 1, 1, 'view')

INSERT INTO public.tbl_module_permissions(id, route_name, display_name, description, module_id, created_by, created_on, full_access_permission, parent_id, assign_permission,order_index, slug_name) VALUES (12, '/membersgroup/newgroup', 'Create', '', 12, 1, '2023-03-14 11:09:12', 1, 0, 1, 2, 'create')

INSERT INTO public.tbl_module_permissions(id, route_name, display_name, description, module_id, created_by, created_on, full_access_permission, parent_id, assign_permission,order_index, slug_name) VALUES (13, '/membersgroup/updategroup', 'Update', '', 12, 1, '2023-03-14 11:09:12', 1, 0, 1, 3, 'update')

INSERT INTO public.tbl_module_permissions(id, route_name, display_name, description, module_id, created_by, created_on, full_access_permission, parent_id, assign_permission,order_index, slug_name) VALUES (14, '/membersgroup/deletegroup', 'Delete', '', 12, 1, '2023-03-14 11:09:12', 1, 0, 1, 4, 'delete')

INSERT INTO public.tbl_module_permissions(id, route_name, display_name, description, module_id, created_by, created_on, full_access_permission, parent_id, assign_permission,order_index, slug_name) VALUES (15, '/member/', 'View', '', 11, 1, '2023-03-14 11:09:12', 1, 0, 1, 1, 'view')

INSERT INTO public.tbl_module_permissions(id, route_name, display_name, description, module_id, created_by, created_on, full_access_permission, parent_id, assign_permission,order_index, slug_name) VALUES (16, '/member/newmember', 'Create', '', 11, 1, '2023-03-14 11:09:12', 1, 0, 1, 2, 'create')

INSERT INTO public.tbl_module_permissions(id, route_name, display_name, description, module_id, created_by, created_on, full_access_permission, parent_id, assign_permission,order_index, slug_name) VALUES (17, '/member/updatemember', 'Update', '', 11, 1, '2023-03-14 11:09:12', 1, 0, 1, 3, 'update')

INSERT INTO public.tbl_module_permissions(id, route_name, display_name, description, module_id, created_by, created_on, full_access_permission, parent_id, assign_permission,order_index, slug_name) VALUES (18, '/member/deletemember', 'Delete', '', 11, 1, '2023-03-14 11:09:12', 1, 0, 1, 4, 'delete')

INSERT INTO public.tbl_module_permissions(id, route_name, display_name, description, module_id, created_by, created_on, full_access_permission, parent_id, assign_permission,order_index, slug_name) VALUES (19, '/memberaccess/', 'Member-Restrict', 'Create and manage content consumption access and restrictions to the Members on the website. ', 13, 1, '2023-03-14 11:09:12', 1, 0, 1, 1, 'memberaccess')

INSERT INTO public.tbl_module_permissions(id, route_name, display_name, description, module_id, created_by, created_on, full_access_permission, parent_id, assign_permission,order_index, slug_name) VALUES (20, '/channel/entrylist/', 'Entries', '', 8, 1, '2023-03-14 11:09:12', 1, 0, 1, 1, 'entries')

INSERT INTO public.tbl_module_permissions(id, route_name, display_name, description, module_id, created_by, created_on, full_access_permission, parent_id, assign_permission,order_index, slug_name) VALUES (21, '/media/', 'View', '', 23, 1, '2023-03-14 11:09:12', 1, 0, 1, 1, 'view')

INSERT INTO public.tbl_module_permissions(id, route_name, display_name, description, module_id, created_by, created_on, full_access_permission, parent_id, assign_permission,order_index, slug_name) VALUES (22, '/settings/roles/rolelist', 'Roles & Permission', 'Create and manage additional roles based on the needs for their tasks and activities in the control panel.', 17, 1, '2023-03-14 11:09:12', 1, 0, 1, 1, 'entries')

INSERT INTO public.tbl_module_permissions(id, route_name, display_name, description, module_id, created_by, created_on, full_access_permission, parent_id, assign_permission,order_index, slug_name) VALUES (23, '/settings/users/userlist', 'Teams', 'Add users for assigning them roles, followed by assigning permissions for appropriate access to the CMS.', 18, 1, '2023-03-14 11:09:12', 1, 0, 1, 1, 'teams')

INSERT INTO public.tbl_module_permissions(id, route_name, display_name, description, module_id, created_by, created_on, full_access_permission, parent_id, assign_permission,order_index, slug_name) VALUES (24, '/settings/languages/languagelist', 'Languages', 'Add different languages here, for the Members to consume content in the language of preference on the website.', 19, 1, '2023-03-14 11:09:12', 1, 0, 1, 1, 'language')

INSERT INTO public.tbl_module_permissions(id, route_name, display_name, description, module_id, created_by, created_on, full_access_permission, parent_id, assign_permission,order_index, slug_name) VALUES (25, '/settings/emails', 'Languages', 'Add different languages here, for the Members to consume content in the language of preference on the website.', 20, 1, '2023-03-14 11:09:12', 1, 0, 1, 1, 'email-templates')

INSERT INTO public.tbl_module_permissions(id, route_name, display_name, description, module_id, created_by, created_on, full_access_permission, parent_id, assign_permission,order_index, slug_name) VALUES (26, '/settings/data', 'Data', 'Add different languages here, for the Members to consume content in the language of preference on the website.', 21, 1, '2023-03-14 11:09:12', 1, 0, 1, 1, 'data')