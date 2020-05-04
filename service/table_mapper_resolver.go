//table 映射器
package service

import (
	"dc/models"
)
//table - model的映射

var	TableResolverToModel = map[string]func(multiModels []interface{}, length int){
		tablePrefix+"thread": ResolverThread,
		tablePrefix+"thread_roadbook": ResolverThreadRoadbook,
		tablePrefix+"2handcar_apply": Resolver2handcarApply,
		tablePrefix+"2handcar_booth":Resolver2handcarBooth,
		tablePrefix+"2handcar_complain":Resolver2handcarComplain,
		tablePrefix+"2handcar_config":Resolver2handcarConfig,
		tablePrefix+"2handcar_debit_list":Resolver2handcarDebitList,
		tablePrefix+"2handcar_intention":Resolver2handcarIntention,
		tablePrefix+"2handcar_list":Resolver2handcarList,
		tablePrefix+"2handcar_order":Resolver2handcarOrder,
		tablePrefix+"2handcar_park":Resolver2handcarPark,
		tablePrefix+"2handcar_park_time":Resolver2handcarParkTime,
		tablePrefix+"2handcar_park_type":Resolver2handcarParkType,
		tablePrefix+"2handcar_plan":Resolver2handcarPlan,
		tablePrefix+"2handcar_time":Resolver2handcarTime,
		tablePrefix+"2handcar_wx_fail":Resolver2handcarWxFail,
		tablePrefix+"account_manage":ResolverAccountManage,
		tablePrefix+"action_log":ResolverActionLog,
		tablePrefix+"activity_area":ResolverActivityArea,
		tablePrefix+"activity_forward":ResolverActivityForward,
		tablePrefix+"activity_forward_content":ResolverActivityForwardContent,
		tablePrefix+"activity_forward_prize":ResolverActivityForwardPrize,
		tablePrefix+"activity_forward_record":ResolverActivityForwardRecord,
		tablePrefix+"activity_forward_record_details":ResolverActivityForwardRecordDetails,
		tablePrefix+"activity_line":ResolverActivityLine,
		tablePrefix+"activity_line_node_related":ResolverActivityLineNodeRelated,
		tablePrefix+"activity_lottery_attendance":ResolverActivityLotteryAttendance,
		tablePrefix+"activity_lottery_auth":ResolverActivityLotteryAuth,
		tablePrefix+"activity_lottery_code":ResolverActivityLotteryCode,
		tablePrefix+"activity_lottery_prize":ResolverActivityLotteryPrize,
		tablePrefix+"activity_lottery_rotary":ResolverActivityLotteryRotary,
		tablePrefix+"activity_lottery_sms":ResolverActivityLotterySms,
		tablePrefix+"activity_lottery_status":ResolverActivityLotteryStatus,
		tablePrefix+"activity_table":ResolverActivityTable,
		tablePrefix+"activity_tags":ResolverActivityTags,
		tablePrefix+"activity_turntable":ResolverActivityTurntable,
		tablePrefix+"activity_turntable_content":ResolverActivityTurntableContent,
		tablePrefix+"activity_turntable_prize":ResolverActivityTurntablePrize,
		tablePrefix+"activity_turntable_qrcodetag":ResolverActivityTurntableQrcodetag,
		tablePrefix+"activity_turntable_record":ResolverActivityTurntableRecord,
		tablePrefix+"activity_turntable_result":ResolverActivityTurntableResult,
		tablePrefix+"activity_turntable_saddress":ResolverActivityTurntableSaddress,
		tablePrefix+"activity_turntable_share":ResolverActivityTurntableShare,
		tablePrefix+"activity_turntable_user":ResolverActivityTurntableUser,
		tablePrefix+"activity_vote_refer_opt":ResolverActivityVoteReferOpt,
		tablePrefix+"activity_vote_rules":ResolverActivityVoteRules,
		tablePrefix+"activity_vote_user_limit":ResolverActivityVoteUserLimit,
		tablePrefix+"activity_vote_user_table":ResolverActivityVoteUserTable,
		tablePrefix+"activity_vote_user_ticketcert":ResolverActivityVoteUserTicketcert,
		tablePrefix+"admin_auth_project_xc":ResolverAdminAuthProjectXc,
		tablePrefix+"admin_auth_project_xc_copy":ResolverAdminAuthProjectXcCopy,
		tablePrefix+"admin_auth_user":ResolverAdminAuthUser,
		tablePrefix+"admin_auth_user_copy":ResolverAdminAuthUserCopy,
		tablePrefix+"admin_menu":ResolverAdminMenu,
		tablePrefix+"admin_report":ResolverAdminReport,
		tablePrefix+"admin_role":ResolverAdminRole,
		tablePrefix+"admin_user":ResolverAdminUser,
		tablePrefix+"admin_user2area":ResolverAdminUser2area,
		tablePrefix+"admin_user2area_org":ResolverAdminUser2areaOrg,
		tablePrefix+"admin_user2org":ResolverAdminUser2org,
		tablePrefix+"admin_user2org_copy":ResolverAdminUser2orgCopy,
		tablePrefix+"admin_user2org_setting":ResolverAdminUser2orgSetting,
		tablePrefix+"admin_user_alltags":ResolverAdminUserAlltags,
		tablePrefix+"admin_user_auth":ResolverAdminUserAuth,
		tablePrefix+"admin_user_auth_copy":ResolverAdminUserAuthCopy,
		tablePrefix+"admin_user_avatar":ResolverAdminUserAvatar,
		tablePrefix+"admin_user_copy1":ResolverAdminUserCopy1,
		tablePrefix+"admin_user_department":ResolverAdminUserDepartment,
		tablePrefix+"admin_user_doyen":ResolverAdminUserDoyen,
		tablePrefix+"admin_user_favor":ResolverAdminUserFavor,
		tablePrefix+"admin_user_impression":ResolverAdminUserImpression,
		tablePrefix+"admin_user_lock":ResolverAdminUserLock,
		tablePrefix+"admin_user_login_fail":ResolverAdminUserLoginFail,
		tablePrefix+"admin_user_money":ResolverAdminUserMoney,
		tablePrefix+"admin_user_money_bill":ResolverAdminUserMoneyBill,
		tablePrefix+"admin_user_money_recharge":ResolverAdminUserMoneyRecharge,
		tablePrefix+"admin_user_money_withdraw":ResolverAdminUserMoneyWithdraw,
		tablePrefix+"admin_user_org":ResolverAdminUserOrg,
		tablePrefix+"admin_user_position":ResolverAdminUserPosition,
		tablePrefix+"admin_user_punish":ResolverAdminUserPunish,
		tablePrefix+"admin_user_real_name_auth":ResolverAdminUserRealNameAuth,
		tablePrefix+"admin_user_role":ResolverAdminUserRole,
		tablePrefix+"admin_user_tags":ResolverAdminUserTags,
		tablePrefix+"admin_user_type":ResolverAdminUserType,
		tablePrefix+"advertisement_admin_log":ResolverAdvertisementAdminLog,
		tablePrefix+"advertisement_content":ResolverAdvertisementContent,
		tablePrefix+"advertisement_org_strategy":ResolverAdvertisementOrgStrategy,
		tablePrefix+"advertisement_position":ResolverAdvertisementPosition,
		tablePrefix+"agreement":ResolverAgreement,
		tablePrefix+"agreement_place":ResolverAgreementPlace,
		tablePrefix+"album":ResolverAlbum,
		tablePrefix+"alltags":ResolverAlltags,
		tablePrefix+"alltags_article":ResolverAlltagsArticle,
		tablePrefix+"alltags_category":ResolverAlltagsCategory,
		tablePrefix+"alltags_category_activity":ResolverAlltagsCategoryActivity,
		tablePrefix+"area":ResolverArea,
		tablePrefix+"area_bak":ResolverAreaBak,
		tablePrefix+"area_hot":ResolverAreaHot,
		tablePrefix+"article":ResolverArticle,
		tablePrefix+"article_category":ResolverArticleCategory,
		tablePrefix+"article_data":ResolverArticleData,
		tablePrefix+"article_tag":ResolverArticleTag,
		tablePrefix+"ask_cate":ResolverAskCate,
		tablePrefix+"ask_expert":ResolverAskExpert,
		tablePrefix+"async_task":ResolverAsyncTask,
		tablePrefix+"attachment":ResolverAttachment,
		tablePrefix+"auth_action":ResolverAuthAction,
		tablePrefix+"auth_role":ResolverAuthRole,
		tablePrefix+"auth_role_action":ResolverAuthRoleAction,
		tablePrefix+"auth_role_source":ResolverAuthRoleSource,
		tablePrefix+"auth_user_role":ResolverAuthUserRole,
		tablePrefix+"baidu_config":ResolverBaiduConfig,
		tablePrefix+"baidu_notice":ResolverBaiduNotice,
		tablePrefix+"baidu_share":ResolverBaiduShare,
		tablePrefix+"benefit":ResolverBenefit,
		tablePrefix+"benefit_access":ResolverBenefitAccess,
		tablePrefix+"benefit_store":ResolverBenefitStore,
		tablePrefix+"bid_keyword":ResolverBidKeyword,
		tablePrefix+"brand":ResolverBrand,
		tablePrefix+"browse_history":ResolverBrowseHistory,
		tablePrefix+"call_audience":ResolverCallAudience,
		tablePrefix+"call_audience_access":ResolverCallAudienceAccess,
		tablePrefix+"call_audience_conf":ResolverCallAudienceConf,
		tablePrefix+"call_audience_conf_val":ResolverCallAudienceConfVal,
		tablePrefix+"call_audience_review":ResolverCallAudienceReview,
		tablePrefix+"call_audience_review_log":ResolverCallAudienceReviewLog,
		tablePrefix+"call_audience_smstask":ResolverCallAudienceSmstask,
		tablePrefix+"call_audience_smstask_content":ResolverCallAudienceSmstaskContent,
		tablePrefix+"call_audience_smstask_record":ResolverCallAudienceSmstaskRecord,
		tablePrefix+"call_audience_user":ResolverCallAudienceUser,
		tablePrefix+"call_audience_visit":ResolverCallAudienceVisit,
		tablePrefix+"call_datatype":ResolverCallDatatype,
		tablePrefix+"call_draw":ResolverCallDraw,
		tablePrefix+"call_event_auth":ResolverCallEventAuth,
		tablePrefix+"call_privacy":ResolverCallPrivacy,
		tablePrefix+"call_questionnaire":ResolverCallQuestionnaire,
		tablePrefix+"call_questionnaire_conf":ResolverCallQuestionnaireConf,
		tablePrefix+"call_questionnaire_user":ResolverCallQuestionnaireUser,
		tablePrefix+"call_questionnaire_val":ResolverCallQuestionnaireVal,
		tablePrefix+"call_record":ResolverCallRecord,
		tablePrefix+"call_user_auth":ResolverCallUserAuth,
		tablePrefix+"car":ResolverCar,
		tablePrefix+"car2conf":ResolverCar2conf,
		tablePrefix+"car_color":ResolverCarColor,
		tablePrefix+"car_conf":ResolverCarConf,
		tablePrefix+"car_conf_val":ResolverCarConfVal,
		tablePrefix+"car_copy":ResolverCarCopy,
		tablePrefix+"car_custom_made":ResolverCarCustomMade,
		tablePrefix+"car_find":ResolverCarFind,
		tablePrefix+"car_love":ResolverCarLove,
		tablePrefix+"car_order":ResolverCarOrder,
		tablePrefix+"car_order_list":ResolverCarOrderList,
		tablePrefix+"car_order_msg":ResolverCarOrderMsg,
		tablePrefix+"car_order_offer":ResolverCarOrderOffer,
		tablePrefix+"car_order_offer_count":ResolverCarOrderOfferCount,
		tablePrefix+"car_order_push_list":ResolverCarOrderPushList,
		tablePrefix+"car_order_tags":ResolverCarOrderTags,
		tablePrefix+"car_pic":ResolverCarPic,
		tablePrefix+"car_search_log":ResolverCarSearchLog,
		tablePrefix+"car_series":ResolverCarSeries,
		tablePrefix+"car_series_color":ResolverCarSeriesColor,
		tablePrefix+"carorder":ResolverCarorder,
		tablePrefix+"carorder_list":ResolverCarorderList,
		tablePrefix+"carorder_lovecar":ResolverCarorderLovecar,
		tablePrefix+"carorder_offer":ResolverCarorderOffer,
		tablePrefix+"carorder_offer_record":ResolverCarorderOfferRecord,
		tablePrefix+"carorder_push_list":ResolverCarorderPushList,
		tablePrefix+"carorder_record":ResolverCarorderRecord,
		tablePrefix+"carorder_tactics":ResolverCarorderTactics,
		tablePrefix+"carorder_tactics_config":ResolverCarorderTacticsConfig,
		tablePrefix+"carorder_tactics_gift":ResolverCarorderTacticsGift,
		tablePrefix+"carorder_tactics_record":ResolverCarorderTacticsRecord,
		tablePrefix+"carousel":ResolverCarousel,
		tablePrefix+"carousel_item":ResolverCarouselItem,
		tablePrefix+"carousel_object":ResolverCarouselObject,
		tablePrefix+"carousel_org_plan":ResolverCarouselOrgPlan,
		tablePrefix+"carousel_org_right":ResolverCarouselOrgRight,
		tablePrefix+"carousel_originality":ResolverCarouselOriginality,
		tablePrefix+"carousel_scheduling":ResolverCarouselScheduling,
		tablePrefix+"carousel_scheduling_authorize":ResolverCarouselSchedulingAuthorize,
		tablePrefix+"carousel_strategy":ResolverCarouselStrategy,
		tablePrefix+"carousel_subject_item":ResolverCarouselSubjectItem,
		tablePrefix+"cert":ResolverCert,
		tablePrefix+"cert_apply":ResolverCertApply,
		tablePrefix+"cert_content":ResolverCertContent,
		tablePrefix+"cert_handout":ResolverCertHandout,
		tablePrefix+"cert_nums_record":ResolverCertNumsRecord,
		tablePrefix+"cert_user":ResolverCertUser,
		tablePrefix+"cert_user_copy":ResolverCertUserCopy,
		tablePrefix+"cert_user_record":ResolverCertUserRecord,
		tablePrefix+"cert_user_record_act":ResolverCertUserRecordAct,
		tablePrefix+"chicken_soup":ResolverChickenSoup,
		tablePrefix+"collect":ResolverCollect,
		tablePrefix+"collect_mapping":ResolverCollectMapping,
		tablePrefix+"comment":ResolverComment,
		tablePrefix+"comment_log":ResolverCommentLog,
		tablePrefix+"config":ResolverConfig,
		tablePrefix+"country":ResolverCountry,
		tablePrefix+"coupon":ResolverCoupon,
		tablePrefix+"coupon_car":ResolverCouponCar,
		tablePrefix+"coupon_category":ResolverCouponCategory,
		tablePrefix+"coupon_check":ResolverCouponCheck,
		tablePrefix+"coupon_code":ResolverCouponCode,
		tablePrefix+"coupon_dealers":ResolverCouponDealers,
		tablePrefix+"coupon_distribute":ResolverCouponDistribute,
		tablePrefix+"coupon_distribute_list":ResolverCouponDistributeList,
		tablePrefix+"coupon_distribute_list_reissue":ResolverCouponDistributeListReissue,
		tablePrefix+"coupon_distribute_reissue":ResolverCouponDistributeReissue,
		tablePrefix+"coupon_issue":ResolverCouponIssue,
		tablePrefix+"coupon_order":ResolverCouponOrder,
		tablePrefix+"coupon_push":ResolverCouponPush,
		tablePrefix+"coupon_push_log":ResolverCouponPushLog,
		tablePrefix+"coupon_refund":ResolverCouponRefund,
		tablePrefix+"coupon_replace":ResolverCouponReplace,
		tablePrefix+"coupon_replace_distribute":ResolverCouponReplaceDistribute,
		tablePrefix+"coupon_replace_reissue":ResolverCouponReplaceReissue,
		tablePrefix+"coupon_share_log":ResolverCouponShareLog,
		tablePrefix+"coupon_ticket":ResolverCouponTicket,
		tablePrefix+"coupon_warehouse":ResolverCouponWarehouse,
		tablePrefix+"coupon_warehouse_code":ResolverCouponWarehouseCode,
		tablePrefix+"crawler_data":ResolverCrawlerData,
		tablePrefix+"danger":ResolverDanger,
		tablePrefix+"dshow_banner":ResolverDshowBanner,
		tablePrefix+"dshow_base":ResolverDshowBase,
		tablePrefix+"dshow_gift":ResolverDshowGift,
		tablePrefix+"dshow_share":ResolverDshowShare,
		tablePrefix+"dshow_text":ResolverDshowText,
		tablePrefix+"easemob_messages":ResolverEasemobMessages,
		tablePrefix+"easemob_messages_num":ResolverEasemobMessagesNum,
		tablePrefix+"easemob_messages_payload":ResolverEasemobMessagesPayload,
		tablePrefix+"emoji":ResolverEmoji,
		tablePrefix+"emoji_group":ResolverEmojiGroup,
		tablePrefix+"emoji_user_access":ResolverEmojiUserAccess,
		tablePrefix+"entrance_way":ResolverEntranceWay,
		tablePrefix+"event":ResolverEvent,
		tablePrefix+"event_activity":ResolverEventActivity,
		tablePrefix+"event_activity_data":ResolverEventActivityData,
		tablePrefix+"event_ad":ResolverEventAd,
		tablePrefix+"event_bdconfig":ResolverEventBdconfig,
		tablePrefix+"event_booth":ResolverEventBooth,
		tablePrefix+"event_brand":ResolverEventBrand,
		tablePrefix+"event_brand_merge":ResolverEventBrandMerge,
		tablePrefix+"event_business":ResolverEventBusiness,
		tablePrefix+"event_desc_img":ResolverEventDescImg,
		tablePrefix+"event_examine":ResolverEventExamine,
		tablePrefix+"event_guest":ResolverEventGuest,
		tablePrefix+"event_hall":ResolverEventHall,
		tablePrefix+"event_img_config":ResolverEventImgConfig,
		tablePrefix+"event_img_unshow":ResolverEventImgUnshow,
		tablePrefix+"event_information":ResolverEventInformation,
		tablePrefix+"event_model":ResolverEventModel,
		tablePrefix+"event_navigation":ResolverEventNavigation,
		tablePrefix+"event_navigation_conf":ResolverEventNavigationConf,
		tablePrefix+"event_newcar":ResolverEventNewcar,
		tablePrefix+"event_schedule":ResolverEventSchedule,
		tablePrefix+"event_showcar":ResolverEventShowcar,
		tablePrefix+"event_staff":ResolverEventStaff,
		tablePrefix+"event_staff_auth":ResolverEventStaffAuth,
		tablePrefix+"event_staff_auth_area":ResolverEventStaffAuthArea,
		tablePrefix+"event_staff_auth_brand":ResolverEventStaffAuthBrand,
		tablePrefix+"event_staff_auth_user":ResolverEventStaffAuthUser,
		tablePrefix+"event_user":ResolverEventUser,
		tablePrefix+"fab":ResolverFab,
		tablePrefix+"factory":ResolverFactory,
		tablePrefix+"failed_jobs":ResolverFailedJobs,
		tablePrefix+"feedback":ResolverFeedback,
		tablePrefix+"feeds_ignore":ResolverFeedsIgnore,
		tablePrefix+"feeds_topic":ResolverFeedsTopic,
		tablePrefix+"feeds_user":ResolverFeedsUser,
		tablePrefix+"follow_group":ResolverFollowGroup,
		tablePrefix+"friend":ResolverFriend,
		tablePrefix+"friend_buyinfo":ResolverFriendBuyinfo,
		tablePrefix+"friend_group":ResolverFriendGroup,
		tablePrefix+"friend_group_access":ResolverFriendGroupAccess,
		tablePrefix+"friend_latent":ResolverFriendLatent,
		tablePrefix+"friend_manage_log":ResolverFriendManageLog,
		tablePrefix+"friend_org":ResolverFriendOrg,
		tablePrefix+"gallery":ResolverGallery,
		tablePrefix+"group":ResolverGroup,
		tablePrefix+"group_authentication":ResolverGroupAuthentication,
		tablePrefix+"group_blacklist":ResolverGroupBlacklist,
		tablePrefix+"group_category":ResolverGroupCategory,
		tablePrefix+"group_invitation":ResolverGroupInvitation,
		tablePrefix+"group_notification":ResolverGroupNotification,
		tablePrefix+"group_notification_read":ResolverGroupNotificationRead,
		tablePrefix+"group_qrcode":ResolverGroupQrcode,
		tablePrefix+"group_tag":ResolverGroupTag,
		tablePrefix+"group_user":ResolverGroupUser,
		tablePrefix+"group_user_kick_out":ResolverGroupUserKickOut,
		tablePrefix+"help_question":ResolverHelpQuestion,
		tablePrefix+"hot_search_ext":ResolverHotSearchExt,
		tablePrefix+"img":ResolverImg,
		tablePrefix+"img_tag":ResolverImgTag,
		tablePrefix+"img_type":ResolverImgType,
		tablePrefix+"industry":ResolverIndustry,
		tablePrefix+"industry_user":ResolverIndustryUser,
		tablePrefix+"koubei":ResolverKoubei,
		tablePrefix+"koubei_album":ResolverKoubeiAlbum,
		tablePrefix+"koubei_spider":ResolverKoubeiSpider,
		tablePrefix+"latent_autoset":ResolverLatentAutoset,
		tablePrefix+"latent_dispense":ResolverLatentDispense,
		tablePrefix+"line":ResolverLine,
		tablePrefix+"line_node":ResolverLineNode,
		tablePrefix+"line_node_album":ResolverLineNodeAlbum,
		tablePrefix+"line_node_comment":ResolverLineNodeComment,
		tablePrefix+"line_node_comment_minutia":ResolverLineNodeCommentMinutia,
		tablePrefix+"line_node_comment_text":ResolverLineNodeCommentText,
		tablePrefix+"line_node_punch_census":ResolverLineNodePunchCensus,
		tablePrefix+"line_node_related":ResolverLineNodeRelated,
		tablePrefix+"live_message":ResolverLiveMessage,
		tablePrefix+"live_user":ResolverLiveUser,
		tablePrefix+"log_admin_login":ResolverLogAdminLogin,
		tablePrefix+"log_admin_work":ResolverLogAdminWork,
		tablePrefix+"lottery":ResolverLottery,
		tablePrefix+"lottery_auth":ResolverLotteryAuth,
		tablePrefix+"lottery_config":ResolverLotteryConfig,
		tablePrefix+"lottery_cron":ResolverLotteryCron,
		tablePrefix+"lottery_new":ResolverLotteryNew,
		tablePrefix+"lottery_new_auth":ResolverLotteryNewAuth,
		tablePrefix+"lottery_new_code_bill":ResolverLotteryNewCodeBill,
		tablePrefix+"lottery_new_config":ResolverLotteryNewConfig,
		tablePrefix+"lottery_new_cron":ResolverLotteryNewCron,
		tablePrefix+"lottery_new_lucky_code":ResolverLotteryNewLuckyCode,
		tablePrefix+"lottery_new_order":ResolverLotteryNewOrder,
		tablePrefix+"lottery_new_org_code":ResolverLotteryNewOrgCode,
		tablePrefix+"lottery_new_prize":ResolverLotteryNewPrize,
		tablePrefix+"lottery_new_record":ResolverLotteryNewRecord,
		tablePrefix+"lottery_new_set":ResolverLotteryNewSet,
		tablePrefix+"lottery_new_user":ResolverLotteryNewUser,
		tablePrefix+"lottery_new_user_code":ResolverLotteryNewUserCode,
		tablePrefix+"lottery_org_code":ResolverLotteryOrgCode,
		tablePrefix+"lottery_part":ResolverLotteryPart,
		tablePrefix+"lottery_prize":ResolverLotteryPrize,
		tablePrefix+"lottery_record":ResolverLotteryRecord,
		tablePrefix+"lottery_thread":ResolverLotteryThread,
		tablePrefix+"lottery_user":ResolverLotteryUser,
		tablePrefix+"lottery_user_code":ResolverLotteryUserCode,
		tablePrefix+"love_chicken":ResolverLoveChicken,
		tablePrefix+"lover_user":ResolverLoverUser,
		tablePrefix+"main_work":ResolverMainWork,
		tablePrefix+"management_type":ResolverManagementType,
		tablePrefix+"manual_code":ResolverManualCode,
		tablePrefix+"manual_code_operation_record":ResolverManualCodeOperationRecord,
		tablePrefix+"menu":ResolverMenu,
		tablePrefix+"menu_copy":ResolverMenuCopy,
		tablePrefix+"message":ResolverMessage,
		tablePrefix+"message_app":ResolverMessageApp,
		tablePrefix+"message_event":ResolverMessageEvent,
		tablePrefix+"message_event_user":ResolverMessageEventUser,
		tablePrefix+"message_extend":ResolverMessageExtend,
		tablePrefix+"message_img":ResolverMessageImg,
		tablePrefix+"message_interact":ResolverMessageInteract,
		tablePrefix+"message_interact_user":ResolverMessageInteractUser,
		tablePrefix+"message_jike":ResolverMessageJike,
		tablePrefix+"message_jike_user":ResolverMessageJikeUser,
		tablePrefix+"message_label":ResolverMessageLabel,
		tablePrefix+"message_lovecar":ResolverMessageLovecar,
		tablePrefix+"message_lovecar_user":ResolverMessageLovecarUser,
		tablePrefix+"message_record":ResolverMessageRecord,
		tablePrefix+"message_setting":ResolverMessageSetting,
		tablePrefix+"message_store":ResolverMessageStore,
		tablePrefix+"message_store_user":ResolverMessageStoreUser,
		tablePrefix+"message_type":ResolverMessageType,
		tablePrefix+"message_type_content":ResolverMessageTypeContent,
		tablePrefix+"message_type_copy":ResolverMessageTypeCopy,
		tablePrefix+"message_url":ResolverMessageUrl,
		tablePrefix+"msg_broker":ResolverMsgBroker,
		tablePrefix+"music":ResolverMusic,
		tablePrefix+"music1":ResolverMusic1,
		tablePrefix+"music_type":ResolverMusicType,
		tablePrefix+"node_punch_auth":ResolverNodePunchAuth,
		tablePrefix+"node_punch_confirm":ResolverNodePunchConfirm,
		tablePrefix+"onlineservice_mute":ResolverOnlineserviceMute,
		tablePrefix+"onlineservice_order":ResolverOnlineserviceOrder,
		tablePrefix+"operate_activity":ResolverOperateActivity,
		tablePrefix+"operate_activity_check":ResolverOperateActivityCheck,
		tablePrefix+"order":ResolverOrder,
		tablePrefix+"order_main":ResolverOrderMain,
		tablePrefix+"org_bank_card":ResolverOrgBankCard,
		tablePrefix+"org_money":ResolverOrgMoney,
		tablePrefix+"org_money_bill":ResolverOrgMoneyBill,
		tablePrefix+"org_money_recharge":ResolverOrgMoneyRecharge,
		tablePrefix+"org_money_withdraw":ResolverOrgMoneyWithdraw,
		tablePrefix+"organization":ResolverOrganization,
		tablePrefix+"organization_activity":ResolverOrganizationActivity,
		tablePrefix+"organization_car":ResolverOrganizationCar,
		tablePrefix+"organization_check":ResolverOrganizationCheck,
		tablePrefix+"organization_check_info":ResolverOrganizationCheckInfo,
		tablePrefix+"organization_conf":ResolverOrganizationConf,
		tablePrefix+"organization_conf_val":ResolverOrganizationConfVal,
		tablePrefix+"organization_copy":ResolverOrganizationCopy,
		tablePrefix+"organization_copy_copy":ResolverOrganizationCopyCopy,
		tablePrefix+"organization_focus":ResolverOrganizationFocus,
		tablePrefix+"organization_log":ResolverOrganizationLog,
		tablePrefix+"organization_order":ResolverOrganizationOrder,
		tablePrefix+"organization_orderpoints_record":ResolverOrganizationOrderpointsRecord,
		tablePrefix+"organization_privilege":ResolverOrganizationPrivilege,
		tablePrefix+"organization_privilege_shop":ResolverOrganizationPrivilegeShop,
		tablePrefix+"organization_privilege_store":ResolverOrganizationPrivilegeStore,
		tablePrefix+"organization_time":ResolverOrganizationTime,
		tablePrefix+"organization_type":ResolverOrganizationType,
		tablePrefix+"partner":ResolverPartner,
		tablePrefix+"partner_blacklist":ResolverPartnerBlacklist,
		tablePrefix+"partner_invite":ResolverPartnerInvite,
		tablePrefix+"partner_log":ResolverPartnerLog,
		tablePrefix+"partner_period":ResolverPartnerPeriod,
		tablePrefix+"partner_qrcode":ResolverPartnerQrcode,
		tablePrefix+"partner_register":ResolverPartnerRegister,
		tablePrefix+"partner_results":ResolverPartnerResults,
		tablePrefix+"partner_task":ResolverPartnerTask,
		tablePrefix+"pay_alipay_log":ResolverPayAlipayLog,
		tablePrefix+"pay_wx_log":ResolverPayWxLog,
		tablePrefix+"phone":ResolverPhone,
		tablePrefix+"photo":ResolverPhoto,
		tablePrefix+"points_data":ResolverPointsData,
		tablePrefix+"points_exceed":ResolverPointsExceed,
		tablePrefix+"points_exp_records":ResolverPointsExpRecords,
		tablePrefix+"points_exp_rules":ResolverPointsExpRules,
		tablePrefix+"points_goods":ResolverPointsGoods,
		tablePrefix+"points_goods_pay":ResolverPointsGoodsPay,
		tablePrefix+"points_log":ResolverPointsLog,
		tablePrefix+"points_product":ResolverPointsProduct,
		tablePrefix+"points_product_records":ResolverPointsProductRecords,
		tablePrefix+"points_ranking":ResolverPointsRanking,
		tablePrefix+"points_sale_rules":ResolverPointsSaleRules,
		tablePrefix+"praise":ResolverPraise,
		tablePrefix+"project_examine":ResolverProjectExamine,
		tablePrefix+"punch":ResolverPunch,
		tablePrefix+"purchase":ResolverPurchase,
		tablePrefix+"purchase_factory":ResolverPurchaseFactory,
		tablePrefix+"purchase_fail":ResolverPurchaseFail,
		tablePrefix+"purchase_img":ResolverPurchaseImg,
		tablePrefix+"purchase_records":ResolverPurchaseRecords,
		tablePrefix+"purchase_sms":ResolverPurchaseSms,
		tablePrefix+"purchase_user":ResolverPurchaseUser,
		tablePrefix+"push_sms_config":ResolverPushSmsConfig,
		tablePrefix+"questionnaire":ResolverQuestionnaire,
		tablePrefix+"questionnaire_answer":ResolverQuestionnaireAnswer,
		tablePrefix+"questionnaire_main":ResolverQuestionnaireMain,
		tablePrefix+"questionnaire_user":ResolverQuestionnaireUser,
		tablePrefix+"rabbit_msg":ResolverRabbitMsg,
		tablePrefix+"recommend":ResolverRecommend,
		tablePrefix+"recommend_auth":ResolverRecommendAuth,
		tablePrefix+"recommend_content":ResolverRecommendContent,
		tablePrefix+"recommend_detail_201905":ResolverRecommendDetail201905,
		tablePrefix+"recommend_detail_201906":ResolverRecommendDetail201906,
		tablePrefix+"recommend_detail_201907":ResolverRecommendDetail201907,
		tablePrefix+"recommend_detail_201908":ResolverRecommendDetail201908,
		tablePrefix+"recommend_detail_201909":ResolverRecommendDetail201909,
		tablePrefix+"recommend_detail_201910":ResolverRecommendDetail201910,
		tablePrefix+"recommend_detail_201911":ResolverRecommendDetail201911,
		tablePrefix+"recommend_detail_201912":ResolverRecommendDetail201912,
		tablePrefix+"recommend_org":ResolverRecommendOrg,
		tablePrefix+"recommend_position":ResolverRecommendPosition,
		tablePrefix+"recommend_position_config":ResolverRecommendPositionConfig,
		tablePrefix+"recommend_record":ResolverRecommendRecord,
		tablePrefix+"recommend_time":ResolverRecommendTime,
		tablePrefix+"recommend_time_extend":ResolverRecommendTimeExtend,
		tablePrefix+"recommend_tongji_day":ResolverRecommendTongjiDay,
		tablePrefix+"recommend_tongji_month":ResolverRecommendTongjiMonth,
		tablePrefix+"redpack":ResolverRedpack,
		tablePrefix+"redpack_area":ResolverRedpackArea,
		tablePrefix+"redpack_check_record":ResolverRedpackCheckRecord,
		tablePrefix+"redpack_content":ResolverRedpackContent,
		tablePrefix+"redpack_msg_record":ResolverRedpackMsgRecord,
		tablePrefix+"redpack_org":ResolverRedpackOrg,
		tablePrefix+"redpack_position":ResolverRedpackPosition,
		tablePrefix+"redpack_prize":ResolverRedpackPrize,
		tablePrefix+"redpack_record":ResolverRedpackRecord,
		tablePrefix+"redpack_record_temp":ResolverRedpackRecordTemp,
		tablePrefix+"redpack_time":ResolverRedpackTime,
		tablePrefix+"redpack_user_card":ResolverRedpackUserCard,
		tablePrefix+"redpack_user_card_record":ResolverRedpackUserCardRecord,
		tablePrefix+"region":ResolverRegion,
		tablePrefix+"report":ResolverReport,
		tablePrefix+"report_img":ResolverReportImg,
		tablePrefix+"report_record":ResolverReportRecord,
		tablePrefix+"ride":ResolverRide,
		tablePrefix+"ride_check":ResolverRideCheck,
		tablePrefix+"ride_check_notes":ResolverRideCheckNotes,
		tablePrefix+"ride_line":ResolverRideLine,
		tablePrefix+"ride_node":ResolverRideNode,
		tablePrefix+"road_book":ResolverRoadBook,
		tablePrefix+"road_book_line":ResolverRoadBookLine,
		tablePrefix+"road_book_line_content":ResolverRoadBookLineContent,
		tablePrefix+"road_book_meet":ResolverRoadBookMeet,
		tablePrefix+"road_book_recommend":ResolverRoadBookRecommend,
		tablePrefix+"roadbook_arouse":ResolverRoadbookArouse,
		tablePrefix+"roadbook_arouse_abnormal":ResolverRoadbookArouseAbnormal,
		tablePrefix+"roadbook_arouse_account":ResolverRoadbookArouseAccount,
		tablePrefix+"roadbook_arouse_attend":ResolverRoadbookArouseAttend,
		tablePrefix+"roadbook_arouse_attend_census":ResolverRoadbookArouseAttendCensus,
		tablePrefix+"roadbook_arouse_flow":ResolverRoadbookArouseFlow,
		tablePrefix+"roadbook_arouse_log":ResolverRoadbookArouseLog,
		tablePrefix+"roadbook_arouse_real_account":ResolverRoadbookArouseRealAccount,
		tablePrefix+"roadbook_arouse_related":ResolverRoadbookArouseRelated,
		tablePrefix+"roadbook_line":ResolverRoadbookLine,
		tablePrefix+"roadbook_merchant":ResolverRoadbookMerchant,
		tablePrefix+"roadbook_merchant_apply":ResolverRoadbookMerchantApply,
		tablePrefix+"roadbook_merchant_invite":ResolverRoadbookMerchantInvite,
		tablePrefix+"roadbook_node_punch_census":ResolverRoadbookNodePunchCensus,
		tablePrefix+"roulette_lottery":ResolverRouletteLottery,
		tablePrefix+"roulette_lottery_attend":ResolverRouletteLotteryAttend,
		tablePrefix+"roulette_lottery_award":ResolverRouletteLotteryAward,
		tablePrefix+"saleinfo_org":ResolverSaleinfoOrg,
		tablePrefix+"saler_star":ResolverSalerStar,
		tablePrefix+"search_log":ResolverSearchLog,
		tablePrefix+"series":ResolverSeries,
		tablePrefix+"service_city":ResolverServiceCity,
		tablePrefix+"sms_api_log":ResolverSmsApiLog,
		tablePrefix+"sms_template":ResolverSmsTemplate,
		tablePrefix+"sms_template_log":ResolverSmsTemplateLog,
		tablePrefix+"sms_total":ResolverSmsTotal,
		tablePrefix+"sms_verify_code":ResolverSmsVerifyCode,
		tablePrefix+"special":ResolverSpecial,
		tablePrefix+"spm_app_download_log":ResolverSpmAppDownloadLog,
		tablePrefix+"spm_baidureferer_log":ResolverSpmBaidurefererLog,
		tablePrefix+"spm_cate":ResolverSpmCate,
		tablePrefix+"spm_designation":ResolverSpmDesignation,
		tablePrefix+"spm_keywords":ResolverSpmKeywords,
		tablePrefix+"spm_log":ResolverSpmLog,
		tablePrefix+"spm_log_copy":ResolverSpmLogCopy,
		tablePrefix+"spm_parameter":ResolverSpmParameter,
		tablePrefix+"spm_place":ResolverSpmPlace,
		tablePrefix+"spm_promoter":ResolverSpmPromoter,
		tablePrefix+"staff":ResolverStaff,
		tablePrefix+"statistics":ResolverStatistics,
		tablePrefix+"supertopic_apply":ResolverSupertopicApply,
		tablePrefix+"tag_car":ResolverTagCar,
		tablePrefix+"tag_cart_specific":ResolverTagCartSpecific,
		tablePrefix+"tags":ResolverTags,
		tablePrefix+"tags_category":ResolverTagsCategory,
		tablePrefix+"tagscategory":ResolverTagscategory,
		tablePrefix+"thread":ResolverThread,
		tablePrefix+"thread_activity":ResolverThreadActivity,
		tablePrefix+"thread_activity_apply":ResolverThreadActivityApply,
		tablePrefix+"thread_activity_car":ResolverThreadActivityCar,
		tablePrefix+"thread_activity_master":ResolverThreadActivityMaster,
		tablePrefix+"thread_activity_scene":ResolverThreadActivityScene,
		tablePrefix+"thread_activity_type":ResolverThreadActivityType,
		tablePrefix+"thread_activity_user":ResolverThreadActivityUser,
		tablePrefix+"thread_ask":ResolverThreadAsk,
		tablePrefix+"thread_ask_answer":ResolverThreadAskAnswer,
		tablePrefix+"thread_ask_answer_follow":ResolverThreadAskAnswerFollow,
		tablePrefix+"thread_census_ranking":ResolverThreadCensusRanking,
		tablePrefix+"thread_cycle":ResolverThreadCycle,
		tablePrefix+"thread_delete_reason":ResolverThreadDeleteReason,
		tablePrefix+"thread_evaluate":ResolverThreadEvaluate,
		tablePrefix+"thread_evaluate_car":ResolverThreadEvaluateCar,
		tablePrefix+"thread_goods":ResolverThreadGoods,
		tablePrefix+"thread_goods_category":ResolverThreadGoodsCategory,
		tablePrefix+"thread_integral_log":ResolverThreadIntegralLog,
		tablePrefix+"thread_mastersay":ResolverThreadMastersay,
		tablePrefix+"thread_mastersay_car":ResolverThreadMastersayCar,
		tablePrefix+"thread_mood":ResolverThreadMood,
		tablePrefix+"thread_play":ResolverThreadPlay,
		tablePrefix+"thread_roadbook":ResolverThreadRoadbook,
		tablePrefix+"thread_roadbook_annex":ResolverThreadRoadbookAnnex,
		tablePrefix+"thread_roadbook_line_been":ResolverThreadRoadbookLineBeen,
		tablePrefix+"thread_roadbook_line_related":ResolverThreadRoadbookLineRelated,
		tablePrefix+"thread_roadbook_note_user":ResolverThreadRoadbookNoteUser,
		tablePrefix+"thread_roadbook_punch_census":ResolverThreadRoadbookPunchCensus,
		tablePrefix+"thread_show":ResolverThreadShow,
		tablePrefix+"thread_topic":ResolverThreadTopic,
		tablePrefix+"thread_topic_access":ResolverThreadTopicAccess,
		tablePrefix+"thread_topic_apply":ResolverThreadTopicApply,
		tablePrefix+"thread_topic_fans":ResolverThreadTopicFans,
		tablePrefix+"thread_visible":ResolverThreadVisible,
		tablePrefix+"thread_word_column":ResolverThreadWordColumn,
		tablePrefix+"thread_words":ResolverThreadWords,
		tablePrefix+"ticket":ResolverTicket,
		tablePrefix+"ticket_api_log":ResolverTicketApiLog,
		tablePrefix+"ticket_apply_user":ResolverTicketApplyUser,
		tablePrefix+"ticket_check_record":ResolverTicketCheckRecord,
		tablePrefix+"ticket_checkout_auth":ResolverTicketCheckoutAuth,
		tablePrefix+"ticket_free":ResolverTicketFree,
		tablePrefix+"ticket_free_handout":ResolverTicketFreeHandout,
		tablePrefix+"ticket_nums_record":ResolverTicketNumsRecord,
		tablePrefix+"ticket_paper_record":ResolverTicketPaperRecord,
		tablePrefix+"ticket_qd":ResolverTicketQd,
		tablePrefix+"ticket_qd_sale":ResolverTicketQdSale,
		tablePrefix+"ticket_redpack":ResolverTicketRedpack,
		tablePrefix+"ticket_redpack_log":ResolverTicketRedpackLog,
		tablePrefix+"ticket_refund":ResolverTicketRefund,
		tablePrefix+"ticket_refund_log":ResolverTicketRefundLog,
		tablePrefix+"ticket_send_record":ResolverTicketSendRecord,
		tablePrefix+"ticket_third_party":ResolverTicketThirdParty,
		tablePrefix+"ticket_user":ResolverTicketUser,
		tablePrefix+"tmp_data":ResolverTmpData,
		tablePrefix+"topic_category":ResolverTopicCategory,
		tablePrefix+"topic_exp_log":ResolverTopicExpLog,
		tablePrefix+"topic_log":ResolverTopicLog,
		tablePrefix+"topic_scompere":ResolverTopicScompere,
		tablePrefix+"topic_signin":ResolverTopicSignin,
		tablePrefix+"traffic_statistics":ResolverTrafficStatistics,
		tablePrefix+"traffic_statistics_log":ResolverTrafficStatisticsLog,
		tablePrefix+"traffic_statistics_log_group":ResolverTrafficStatisticsLogGroup,
		tablePrefix+"twitter":ResolverTwitter,
		tablePrefix+"twitter_album":ResolverTwitterAlbum,
		tablePrefix+"twitter_source":ResolverTwitterSource,
		tablePrefix+"twitter_topic_access":ResolverTwitterTopicAccess,
		tablePrefix+"twitter_visibility":ResolverTwitterVisibility,
		tablePrefix+"upload_apk":ResolverUploadApk,
		tablePrefix+"user_bank_card":ResolverUserBankCard,
		tablePrefix+"user_behaviour":ResolverUserBehaviour,
		tablePrefix+"user_car":ResolverUserCar,
		tablePrefix+"user_follow":ResolverUserFollow,
		tablePrefix+"user_points":ResolverUserPoints,
		tablePrefix+"user_recommend":ResolverUserRecommend,
		tablePrefix+"user_setting":ResolverUserSetting,
		tablePrefix+"video":ResolverVideo,
		tablePrefix+"video_attach":ResolverVideoAttach,
		tablePrefix+"video_attach_receive":ResolverVideoAttachReceive,
		tablePrefix+"video_log":ResolverVideoLog,
		tablePrefix+"video_notice":ResolverVideoNotice,
		tablePrefix+"video_recommend_table":ResolverVideoRecommendTable,
		tablePrefix+"video_visible":ResolverVideoVisible,
		tablePrefix+"virtual_nickname":ResolverVirtualNickname,
		tablePrefix+"virtual_record":ResolverVirtualRecord,
		tablePrefix+"wanna_buy":ResolverWannaBuy,
		tablePrefix+"weather_citys":ResolverWeatherCitys,
		tablePrefix+"weather_today":ResolverWeatherToday,
		tablePrefix+"weixin_config":ResolverWeixinConfig,
		tablePrefix+"weixin_custom_send":ResolverWeixinCustomSend,
		tablePrefix+"weixin_kerword":ResolverWeixinKerword,
		tablePrefix+"weixin_material":ResolverWeixinMaterial,
		tablePrefix+"weixin_menu":ResolverWeixinMenu,
		tablePrefix+"weixin_push":ResolverWeixinPush,
		tablePrefix+"weixin_push_auto":ResolverWeixinPushAuto,
		tablePrefix+"weixin_push_config":ResolverWeixinPushConfig,
		tablePrefix+"weixin_push_interactive":ResolverWeixinPushInteractive,
		tablePrefix+"weixin_push_material":ResolverWeixinPushMaterial,
		tablePrefix+"weixin_push_record":ResolverWeixinPushRecord,
		tablePrefix+"weixin_push_tag":ResolverWeixinPushTag,
		tablePrefix+"weixin_qrcode":ResolverWeixinQrcode,
		tablePrefix+"weixin_qrcode_scan_log":ResolverWeixinQrcodeScanLog,
		tablePrefix+"weixin_template":ResolverWeixinTemplate,
		tablePrefix+"xc1_jurisdiction":ResolverXc1Jurisdiction,
		tablePrefix+"xc1_manage_object":ResolverXc1ManageObject,
		tablePrefix+"xc1_org2scene":ResolverXc1Org2scene,
		tablePrefix+"xc1_worker2guest":ResolverXc1Worker2guest,
		tablePrefix+"xc_app_message":ResolverXcAppMessage,
		tablePrefix+"xc_app_message_count":ResolverXcAppMessageCount,
		tablePrefix+"xc_apply_item":ResolverXcApplyItem,
		tablePrefix+"xc_apply_type":ResolverXcApplyType,
		tablePrefix+"xc_area":ResolverXcArea,
		tablePrefix+"xc_auth_mapping":ResolverXcAuthMapping,
		tablePrefix+"xc_authorise":ResolverXcAuthorise,
		tablePrefix+"xc_check_course":ResolverXcCheckCourse,
		tablePrefix+"xc_check_table":ResolverXcCheckTable,
		tablePrefix+"xc_checker2object":ResolverXcChecker2object,
		tablePrefix+"xc_friends":ResolverXcFriends,
		tablePrefix+"xc_guest2object":ResolverXcGuest2object,
		tablePrefix+"xc_job_type":ResolverXcJobType,
		tablePrefix+"xc_message":ResolverXcMessage,
		tablePrefix+"xc_message_addtask":ResolverXcMessageAddtask,
		tablePrefix+"xc_message_apply":ResolverXcMessageApply,
		tablePrefix+"xc_message_content":ResolverXcMessageContent,
		tablePrefix+"xc_message_rectification":ResolverXcMessageRectification,
		tablePrefix+"xc_message_setting":ResolverXcMessageSetting,
		tablePrefix+"xc_message_unqualifiedcheck":ResolverXcMessageUnqualifiedcheck,
		tablePrefix+"xc_object":ResolverXcObject,
		tablePrefix+"xc_object2dispose":ResolverXcObject2dispose,
		tablePrefix+"xc_object_organization":ResolverXcObjectOrganization,
		tablePrefix+"xc_object_type":ResolverXcObjectType,
		tablePrefix+"xc_offset_trigger_time":ResolverXcOffsetTriggerTime,
		tablePrefix+"xc_order":ResolverXcOrder,
		tablePrefix+"xc_order_progress":ResolverXcOrderProgress,
		tablePrefix+"xc_organization":ResolverXcOrganization,
		tablePrefix+"xc_pay_log":ResolverXcPayLog,
		tablePrefix+"xc_permission":ResolverXcPermission,
		tablePrefix+"xc_scene":ResolverXcScene,
		tablePrefix+"xc_scene_review":ResolverXcSceneReview,
		tablePrefix+"xc_staff2object":ResolverXcStaff2object,
		tablePrefix+"xc_task":ResolverXcTask,
		tablePrefix+"xc_task_copy":ResolverXcTaskCopy,
		tablePrefix+"xc_task_dis":ResolverXcTaskDis,
		tablePrefix+"xc_task_disimg":ResolverXcTaskDisimg,
		tablePrefix+"xc_task_log":ResolverXcTaskLog,
		tablePrefix+"xc_task_package":ResolverXcTaskPackage,
		tablePrefix+"xc_top_up_log":ResolverXcTopUpLog,
		tablePrefix+"xc_user2scene":ResolverXcUser2scene,
		tablePrefix+"xc_usercoordinate":ResolverXcUsercoordinate,
		tablePrefix+"xc_worker":ResolverXcWorker,
		tablePrefix+"xc_worker2area":ResolverXcWorker2area,
		tablePrefix+"xc_worker_yanghuan_back":ResolverXcWorkerYanghuanBack,
		tablePrefix+"yaoyiyao":ResolverYaoyiyao,
		tablePrefix+"yaoyiyao_auth":ResolverYaoyiyaoAuth,
		tablePrefix+"yaoyiyao_awards":ResolverYaoyiyaoAwards,
		tablePrefix+"yaoyiyao_awards_copy":ResolverYaoyiyaoAwardsCopy,
		tablePrefix+"yaoyiyao_cacheawards":ResolverYaoyiyaoCacheawards,
		tablePrefix+"yaoyiyao_end_log":ResolverYaoyiyaoEndLog,
		tablePrefix+"yaoyiyao_not_win":ResolverYaoyiyaoNotWin,
		tablePrefix+"yaoyiyao_pass_log":ResolverYaoyiyaoPassLog,
		tablePrefix+"yaoyiyao_rank":ResolverYaoyiyaoRank,
		tablePrefix+"yaoyiyao_rotary":ResolverYaoyiyaoRotary,
		tablePrefix+"yaoyiyao_signin":ResolverYaoyiyaoSignin,
		tablePrefix+"yaoyiyao_status":ResolverYaoyiyaoStatus,
	}

//创建数据 解析thread model 参数
func ResolverThread(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgThread)
		data[k] = tmp.Id
	}
	Ret.Data = data
}
//创建数据 解析thread model 参数
func ResolverThreadRoadbook(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgThreadRoadbook)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

//创建数据 解析user model 参数
func Resolver2handcarApply(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.Tg2handcarApply)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func Resolver2handcarBooth(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.Tg2handcarBooth)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func Resolver2handcarComplain(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.Tg2handcarComplain)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func Resolver2handcarConfig(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.Tg2handcarConfig)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func Resolver2handcarDebitList(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.Tg2handcarDebitList)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func Resolver2handcarIntention(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.Tg2handcarIntention)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func Resolver2handcarList(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.Tg2handcarList)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func Resolver2handcarOrder(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.Tg2handcarOrder)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func Resolver2handcarPark(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.Tg2handcarPark)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func Resolver2handcarParkTime(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.Tg2handcarParkTime)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func Resolver2handcarParkType(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.Tg2handcarParkType)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func Resolver2handcarPlan(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.Tg2handcarPlan)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func Resolver2handcarTime(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.Tg2handcarTime)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func Resolver2handcarWxFail(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.Tg2handcarWxFail)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func ResolverAccountManage(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgAccountManage)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func ResolverActionLog(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgActionLog)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func ResolverActivityArea(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgActivityArea)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func ResolverActivityForward(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgActivityForward)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func ResolverActivityForwardContent(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgActivityForwardContent)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func ResolverActivityForwardPrize(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgActivityForwardPrize)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func ResolverActivityForwardRecord(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgActivityForwardRecord)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func ResolverActivityForwardRecordDetails(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgActivityForwardRecordDetails)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func ResolverActivityLine(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgActivityLine)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func ResolverActivityLineNodeRelated(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgActivityLineNodeRelated)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func ResolverActivityLotteryAttendance(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgActivityLotteryAttendance)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func ResolverActivityLotteryAuth(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgActivityLotteryAuth)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func ResolverActivityLotteryCode(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgActivityLotteryCode)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func ResolverActivityLotteryPrize(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgActivityLotteryPrize)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func ResolverActivityLotteryRotary(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgActivityLotteryRotary)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func ResolverActivityLotterySms(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgActivityLotterySms)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func ResolverActivityLotteryStatus(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgActivityLotteryStatus)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func ResolverActivityTable(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgActivityTable)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func ResolverActivityTags(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgActivityTags)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func ResolverActivityTurntable(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgActivityTurntable)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func ResolverActivityTurntableContent(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgActivityTurntableContent)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func ResolverActivityTurntablePrize(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgActivityTurntablePrize)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func ResolverActivityTurntableQrcodetag(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgActivityTurntableQrcodetag)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func ResolverActivityTurntableRecord(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgActivityTurntableRecord)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func ResolverActivityTurntableResult(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgActivityTurntableResult)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func ResolverActivityTurntableSaddress(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgActivityTurntableSaddress)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func ResolverActivityTurntableShare(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgActivityTurntableShare)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func ResolverActivityTurntableUser(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgActivityTurntableUser)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func ResolverActivityVoteReferOpt(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgActivityVoteReferOpt)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func ResolverActivityVoteRules(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgActivityVoteRules)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func ResolverActivityVoteUserLimit(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgActivityVoteUserLimit)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func ResolverActivityVoteUserTable(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgActivityVoteUserTable)
		data[k] = tmp.VoteId
	}
	Ret.Data = data
}

func ResolverActivityVoteUserTicketcert(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgActivityVoteUserTicketcert)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func ResolverAdminAuthProjectXc(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgAdminAuthProjectXc)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func ResolverAdminAuthProjectXcCopy(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgAdminAuthProjectXcCopy)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func ResolverAdminAuthUser(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgAdminAuthUser)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func ResolverAdminAuthUserCopy(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgAdminAuthUserCopy)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func ResolverAdminMenu(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgAdminMenu)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func ResolverAdminReport(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgAdminReport)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func ResolverAdminRole(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgAdminRole)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func ResolverAdminUser(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgAdminUser)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func ResolverAdminUser2area(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgAdminUser2area)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func ResolverAdminUser2areaOrg(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgAdminUser2areaOrg)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func  ResolverAdminUser2org(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgAdminUser2org)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func  ResolverAdminUser2orgCopy(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgAdminUser2orgCopy)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func  ResolverAdminUser2orgSetting(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgAdminUser2orgSetting)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func  ResolverAdminUserAlltags(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgAdminUserAlltags)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func  ResolverAdminUserAuth(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgAdminUserAuth)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func  ResolverAdminUserAuthCopy(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgAdminUserAuthCopy)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func  ResolverAdminUserAvatar(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgAdminUserAvatar)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func  ResolverAdminUserCopy1(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgAdminUserCopy1)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func  ResolverAdminUserDepartment(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgAdminUserDepartment)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func  ResolverAdminUserDoyen(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgAdminUserDoyen)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func  ResolverAdminUserFavor(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgAdminUserFavor)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func  ResolverAdminUserImpression(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgAdminUserImpression)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func  ResolverAdminUserLock(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgAdminUserLock)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func  ResolverAdminUserLoginFail(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgAdminUserLoginFail)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func  ResolverAdminUserMoney(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgAdminUserMoney)
		data[k] = tmp.Uid
	}
	Ret.Data = data
}

func  ResolverAdminUserMoneyBill(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgAdminUserMoneyBill)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func  ResolverAdminUserMoneyRecharge(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgAdminUserMoneyRecharge)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func  ResolverAdminUserMoneyWithdraw(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgAdminUserMoneyWithdraw)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func  ResolverAdminUserOrg(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgAdminUserOrg)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func  ResolverAdminUserPosition(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgAdminUserPosition)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func  ResolverAdminUserPunish(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgAdminUserPunish)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func  ResolverAdminUserRealNameAuth(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgAdminUserRealNameAuth)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func  ResolverAdminUserRole(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgAdminUserRole)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func  ResolverAdminUserTags(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgAdminUserTags)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func  ResolverAdminUserType(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgAdminUserType)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func  ResolverAdvertisementAdminLog(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgAdvertisementAdminLog)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func  ResolverAdvertisementContent(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgAdvertisementContent)
		data[k] = tmp.ContentId
	}
	Ret.Data = data
}

func  ResolverAdvertisementOrgStrategy(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgAdvertisementOrgStrategy)
		data[k] = tmp.StrategyId
	}
	Ret.Data = data
}

func  ResolverAdvertisementPosition(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgAdvertisementPosition)
		data[k] = tmp.AdPositionId
	}
	Ret.Data = data
}

func  ResolverAgreement(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgAgreement)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func  ResolverAgreementPlace(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgAgreementPlace)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func  ResolverAlbum(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgAlbum)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func  ResolverAlltags(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgAlltags)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func  ResolverAlltagsArticle(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgAlltagsArticle)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func ResolverAlltagsCategory(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgAlltagsCategory)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func ResolverAlltagsCategoryActivity(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgAlltagsCategoryActivity)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func  ResolverArea(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgArea)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func  ResolverAreaBak(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgAreaBak)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func  ResolverAreaHot(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgAreaHot)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func ResolverArticle(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgArticle)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func  ResolverArticleCategory(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgArticleCategory)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func  ResolverArticleData(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgArticleData)
		data[k] = tmp.ArticleId
	}
	Ret.Data = data
}

func  ResolverArticleTag(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgArticleTag)
		data[k] = tmp.ArticleId
	}
	Ret.Data = data
}

func  ResolverAskCate(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgAskCate)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func  ResolverAskExpert(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgAskExpert)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func  ResolverAsyncTask(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgAsyncTask)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func  ResolverAttachment(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgAttachment)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func  ResolverAuthAction(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgAuthAction)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func  ResolverAuthRole(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgAuthRole)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func  ResolverAuthRoleAction(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgAuthRoleAction)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func  ResolverAuthRoleSource(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgAuthRoleSource)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func  ResolverAuthUserRole(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgAuthUserRole)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func  ResolverBaiduConfig(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgBaiduConfig)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func  ResolverBaiduNotice(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgBaiduNotice)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func  ResolverBaiduShare(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgBaiduShare)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func  ResolverBenefit(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgBenefit)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func  ResolverBenefitAccess(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgBenefitAccess)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func  ResolverBenefitStore(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgBenefitStore)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func  ResolverBidKeyword(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgBidKeyword)
		data[k] = tmp.Keywordid
	}
	Ret.Data = data
}

func  ResolverBrand(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgBrand)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func  ResolverBrowseHistory(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgBrowseHistory)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func  ResolverCallAudience(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgCallAudience)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func  ResolverCallAudienceAccess(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgCallAudienceAccess)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func  ResolverCallAudienceConf(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgCallAudienceConf)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func  ResolverCallAudienceConfVal(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgCallAudienceConfVal)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func  ResolverCallAudienceReview(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgCallAudienceReview)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func  ResolverCallAudienceReviewLog(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgCallAudienceReviewLog)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func  ResolverCallAudienceSmstask(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgCallAudienceSmstask)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func  ResolverCallAudienceSmstaskContent(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgCallAudienceSmstaskContent)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func  ResolverCallAudienceSmstaskRecord(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgCallAudienceSmstaskRecord)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func  ResolverCallAudienceUser(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgCallAudienceUser)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func  ResolverCallAudienceVisit(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgCallAudienceVisit)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func  ResolverCallDatatype(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgCallDatatype)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func  ResolverCallDraw(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgCallDraw)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func  ResolverCallEventAuth(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgCallEventAuth)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func  ResolverCallPrivacy(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgCallPrivacy)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func  ResolverCallQuestionnaire(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgCallQuestionnaire)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func  ResolverCallQuestionnaireConf(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgCallQuestionnaireConf)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func  ResolverCallQuestionnaireUser(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgCallQuestionnaireUser)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func  ResolverCallQuestionnaireVal(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgCallQuestionnaireVal)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func  ResolverCallRecord(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgCallRecord)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func  ResolverCallUserAuth(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgCallUserAuth)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func  ResolverCar(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgCar)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func  ResolverCar2conf(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgCar2conf)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func  ResolverCarColor(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgCarColor)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func  ResolverCarConf(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgCarConf)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func  ResolverCarConfVal(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgCarConfVal)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func  ResolverCarCopy(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgCarCopy)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func  ResolverCarCustomMade(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgCarCustomMade)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func  ResolverCarFind(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgCarFind)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func  ResolverCarLove(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgCarLove)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func  ResolverCarOrder(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgCarOrder)
		data[k] = tmp.Id
	}
	Ret.Data = data
}


func ResolverCarOrderList(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgCarOrderList)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func  ResolverCarOrderMsg(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgCarOrderMsg)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func  ResolverCarOrderOffer(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgCarOrderOffer)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func  ResolverCarOrderOfferCount(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgCarOrderOfferCount)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func  ResolverCarOrderPushList(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgCarOrderPushList)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func  ResolverCarOrderTags(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgCarOrderTags)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func  ResolverCarPic(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgCarPic)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func  ResolverCarSearchLog(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgCarSearchLog)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func  ResolverCarSeries(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgCarSeries)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func  ResolverCarSeriesColor(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgCarSeriesColor)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func  ResolverCarorder(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgCarorder)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func ResolverCarorderList(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgCarorderList)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func  ResolverCarorderLovecar(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgCarorderLovecar)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func  ResolverCarorderOffer(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgCarorderOffer)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func  ResolverCarorderOfferRecord(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgCarorderOfferRecord)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func  ResolverCarorderPushList(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgCarorderPushList)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func  ResolverCarorderRecord(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgCarorderRecord)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func  ResolverCarorderTactics(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgCarorderTactics)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func  ResolverCarorderTacticsConfig(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgCarorderTacticsConfig)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func  ResolverCarorderTacticsGift(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgCarorderTacticsGift)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func  ResolverCarorderTacticsRecord(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgCarorderTacticsRecord)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func  ResolverCarousel(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgCarousel)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func  ResolverCarouselItem(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgCarouselItem)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func  ResolverCarouselObject(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgCarouselObject)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func  ResolverCarouselOrgPlan(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgCarouselOrgPlan)
		data[k] = tmp.PlanId
	}
	Ret.Data = data
}

func  ResolverCarouselOrgRight(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgCarouselOrgRight)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func  ResolverCarouselOriginality(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgCarouselOriginality)
		data[k] = tmp.OriginalityId
	}
	Ret.Data = data
}

func  ResolverCarouselScheduling(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgCarouselScheduling)
		data[k] = tmp.SchedulingId
	}
	Ret.Data = data
}

func  ResolverCarouselSchedulingAuthorize(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgCarouselSchedulingAuthorize)
		data[k] = tmp.CarouselSchedulingAuthorizeId
	}
	Ret.Data = data
}

func  ResolverCarouselStrategy(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgCarouselStrategy)
		data[k] = tmp.StrategyId
	}
	Ret.Data = data
}

func  ResolverCarouselSubjectItem(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgCarouselSubjectItem)
		data[k] = tmp.SubjectId
	}
	Ret.Data = data
}

func  ResolverCert(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgCert)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func  ResolverCertApply(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgCertApply)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func  ResolverCertContent(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgCertContent)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func  ResolverCertHandout(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgCertHandout)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func  ResolverCertNumsRecord(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgCertNumsRecord)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func  ResolverCertUser(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgCertUser)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func  ResolverCertUserCopy(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgCertUserCopy)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func  ResolverCertUserRecord(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgCertUserRecord)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func  ResolverCertUserRecordAct(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgCertUserRecordAct)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func  ResolverChickenSoup(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgChickenSoup)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func  ResolverCollect(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgCollect)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func  ResolverCollectMapping(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgCollectMapping)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func  ResolverComment(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgComment)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func  ResolverCommentLog(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgCommentLog)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func  ResolverConfig(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgConfig)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func  ResolverCountry(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgCountry)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func  ResolverCoupon(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgCoupon)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func  ResolverCouponCar(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgCouponCar)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func  ResolverCouponCategory(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgCouponCategory)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func  ResolverCouponCheck(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgCouponCheck)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func  ResolverCouponCode(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgCouponCode)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func  ResolverCouponDealers(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgCouponDealers)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func  ResolverCouponDistribute(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgCouponDistribute)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func  ResolverCouponDistributeList(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgCouponDistributeList)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func  ResolverCouponDistributeListReissue(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgCouponDistributeListReissue)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func  ResolverCouponDistributeReissue(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgCouponDistributeReissue)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func  ResolverCouponIssue(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgCouponIssue)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func  ResolverCouponOrder(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgCouponOrder)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func  ResolverCouponPush(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgCouponPush)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func  ResolverCouponPushLog(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgCouponPushLog)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func  ResolverCouponRefund(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgCouponRefund)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func  ResolverCouponReplace(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgCouponReplace)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func  ResolverCouponReplaceDistribute(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgCouponReplaceDistribute)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func  ResolverCouponReplaceReissue(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgCouponReplaceReissue)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func  ResolverCouponShareLog(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgCouponShareLog)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func  ResolverCouponTicket(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgCouponTicket)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func  ResolverCouponWarehouse(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgCouponWarehouse)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func  ResolverCouponWarehouseCode(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgCouponWarehouseCode)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func  ResolverCrawlerData(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgCrawlerData)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func  ResolverDanger(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgDanger)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func  ResolverDshowBanner(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgDshowBanner)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func  ResolverDshowBase(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgDshowBase)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func  ResolverDshowGift(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgDshowGift)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func  ResolverDshowShare(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgDshowShare)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func  ResolverDshowText(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgDshowText)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func  ResolverEasemobMessages(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgEasemobMessages)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func  ResolverEasemobMessagesNum(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgEasemobMessagesNum)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func  ResolverEasemobMessagesPayload(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgEasemobMessagesPayload)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func  ResolverEmoji(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgEmoji)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func  ResolverEmojiGroup(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgEmojiGroup)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func  ResolverEmojiUserAccess(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgEmojiUserAccess)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func  ResolverEntranceWay(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgEntranceWay)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func  ResolverEvent(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgEvent)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func  ResolverEventActivity(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgEventActivity)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func  ResolverEventActivityData(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgEventActivityData)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func  ResolverEventAd(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgEventAd)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func  ResolverEventBdconfig(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgEventBdconfig)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func  ResolverEventBooth(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgEventBooth)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func  ResolverEventBrand(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgEventBrand)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func  ResolverEventBrandMerge(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgEventBrandMerge)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func  ResolverEventBusiness(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgEventBusiness)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func  ResolverEventDescImg(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgEventDescImg)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func  ResolverEventExamine(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgEventExamine)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func  ResolverEventGuest(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgEventGuest)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func  ResolverEventHall(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgEventHall)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func  ResolverEventImgConfig(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgEventImgConfig)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func  ResolverEventImgUnshow(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgEventImgUnshow)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func  ResolverEventInformation(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgEventInformation)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func  ResolverEventModel(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgEventModel)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func  ResolverEventNavigation(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgEventNavigation)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func  ResolverEventNavigationConf(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgEventNavigationConf)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func  ResolverEventNewcar(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgEventNewcar)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func  ResolverEventSchedule(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgEventSchedule)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func  ResolverEventShowcar(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgEventShowcar)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func  ResolverEventStaff(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgEventStaff)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func  ResolverEventStaffAuth(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgEventStaffAuth)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func  ResolverEventStaffAuthArea(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgEventStaffAuthArea)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func  ResolverEventStaffAuthBrand(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgEventStaffAuthBrand)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func  ResolverEventStaffAuthUser(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgEventStaffAuthUser)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func  ResolverEventUser(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgEventUser)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func  ResolverFab(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgFab)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func  ResolverFactory(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgFactory)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func  ResolverFailedJobs(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgFailedJobs)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func  ResolverFeedback(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgFeedback)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func  ResolverFeedsIgnore(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgFeedsIgnore)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func  ResolverFeedsTopic(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgFeedsTopic)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func  ResolverFeedsUser(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgFeedsUser)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func  ResolverFollowGroup(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgFollowGroup)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func  ResolverFriend(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgFriend)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func  ResolverFriendBuyinfo(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgFriendBuyinfo)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func  ResolverFriendGroup(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgFriendGroup)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func  ResolverFriendGroupAccess(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgFriendGroupAccess)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func  ResolverFriendLatent(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgFriendLatent)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func  ResolverFriendManageLog(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgFriendManageLog)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func  ResolverFriendOrg(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgFriendOrg)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func  ResolverGallery(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgGallery)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func  ResolverGroup(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgGroup)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func  ResolverGroupAuthentication(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgGroupAuthentication)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func  ResolverGroupBlacklist(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgGroupBlacklist)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func  ResolverGroupCategory(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgGroupCategory)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func  ResolverGroupInvitation(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgGroupInvitation)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func  ResolverGroupNotification(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgGroupNotification)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func  ResolverGroupNotificationRead(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgGroupNotificationRead)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func  ResolverGroupQrcode(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgGroupQrcode)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func  ResolverGroupTag(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgGroupTag)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func ResolverGroupUser(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgGroupUser)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func  ResolverGroupUserKickOut(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgGroupUserKickOut)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func  ResolverHelpQuestion(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgHelpQuestion)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func ResolverHotSearchExt(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgHotSearchExt)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func  ResolverImg(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgImg)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func  ResolverImgTag(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgImgTag)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func  ResolverImgType(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgImgType)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func  ResolverIndustry(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgIndustry)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func  ResolverIndustryUser(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgIndustryUser)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func ResolverKoubei(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgKoubei)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func  ResolverKoubeiAlbum(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgKoubeiAlbum)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func  ResolverKoubeiSpider(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgKoubeiSpider)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func  ResolverLatentAutoset(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgLatentAutoset)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func  ResolverLatentDispense(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgLatentDispense)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func  ResolverLine(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgLine)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func  ResolverLineNode(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgLineNode)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func  ResolverLineNodeAlbum(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgLineNodeAlbum)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func  ResolverLineNodeComment(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgLineNodeComment)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func  ResolverLineNodeCommentMinutia(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgLineNodeCommentMinutia)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func  ResolverLineNodeCommentText(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgLineNodeCommentText)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func  ResolverLineNodePunchCensus(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgLineNodePunchCensus)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func  ResolverLineNodeRelated(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgLineNodeRelated)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func  ResolverLiveMessage(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgLiveMessage)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func  ResolverLiveUser(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgLiveUser)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func  ResolverLogAdminLogin(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgLogAdminLogin)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func  ResolverLogAdminWork(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgLogAdminWork)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func  ResolverLottery(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgLottery)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func  ResolverLotteryAuth(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgLotteryAuth)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func  ResolverLotteryConfig(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgLotteryConfig)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func  ResolverLotteryCron(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgLotteryCron)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func  ResolverLotteryNew(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgLotteryNew)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func  ResolverLotteryNewAuth(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgLotteryNewAuth)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func  ResolverLotteryNewCodeBill(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgLotteryNewCodeBill)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func  ResolverLotteryNewConfig(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgLotteryNewConfig)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func  ResolverLotteryNewCron(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgLotteryNewCron)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func  ResolverLotteryNewLuckyCode(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgLotteryNewLuckyCode)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func  ResolverLotteryNewOrder(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgLotteryNewOrder)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func  ResolverLotteryNewOrgCode(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgLotteryNewOrgCode)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func  ResolverLotteryNewPrize(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgLotteryNewPrize)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func  ResolverLotteryNewRecord(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgLotteryNewRecord)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func  ResolverLotteryNewSet(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgLotteryNewSet)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func  ResolverLotteryNewUser(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgLotteryNewUser)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func  ResolverLotteryNewUserCode(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgLotteryNewUserCode)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func  ResolverLotteryOrgCode(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgLotteryOrgCode)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func  ResolverLotteryPart(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgLotteryPart)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func  ResolverLotteryPrize(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgLotteryPrize)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func  ResolverLotteryRecord(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgLotteryRecord)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func  ResolverLotteryThread(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgLotteryThread)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func  ResolverLotteryUser(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgLotteryUser)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func  ResolverLotteryUserCode(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgLotteryUserCode)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func  ResolverLoveChicken(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgLoveChicken)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func  ResolverLoverUser(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgLoverUser)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func  ResolverMainWork(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgMainWork)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func  ResolverManagementType(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgManagementType)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func  ResolverManualCode(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgManualCode)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func  ResolverManualCodeOperationRecord(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgManualCodeOperationRecord)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func  ResolverMenu(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgMenu)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func  ResolverMenuCopy(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgMenuCopy)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func  ResolverMessage(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgMessage)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func  ResolverMessageApp(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgMessageApp)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func  ResolverMessageEvent(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgMessageEvent)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func  ResolverMessageEventUser(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgMessageEventUser)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func  ResolverMessageExtend(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgMessageExtend)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func  ResolverMessageImg(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgMessageImg)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func  ResolverMessageInteract(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgMessageInteract)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func  ResolverMessageInteractUser(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgMessageInteractUser)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func ResolverMessageJike(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgMessageJike)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func  ResolverMessageJikeUser(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgMessageJikeUser)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func  ResolverMessageLabel(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgMessageLabel)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func  ResolverMessageLovecar(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgMessageLovecar)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func  ResolverMessageLovecarUser(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgMessageLovecarUser)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func  ResolverMessageRecord(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgMessageRecord)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func  ResolverMessageSetting(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgMessageSetting)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func  ResolverMessageStore(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgMessageStore)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func  ResolverMessageStoreUser(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgMessageStoreUser)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func  ResolverMessageType(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgMessageType)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func  ResolverMessageTypeContent(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgMessageTypeContent)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func  ResolverMessageTypeCopy(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgMessageTypeCopy)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func  ResolverMessageUrl(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgMessageUrl)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func  ResolverMsgBroker(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgMsgBroker)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func  ResolverMusic(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgMusic)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func  ResolverMusic1(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgMusic1)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func  ResolverMusicType(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgMusicType)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func  ResolverNodePunchAuth(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgNodePunchAuth)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func  ResolverNodePunchConfirm(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgNodePunchConfirm)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func  ResolverOnlineserviceMute(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgOnlineserviceMute)
		data[k] = tmp.Uid
	}
	Ret.Data = data
}

func  ResolverOnlineserviceOrder(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgOnlineserviceOrder)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func  ResolverOperateActivity(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgOperateActivity)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func  ResolverOperateActivityCheck(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgOperateActivityCheck)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func  ResolverOrder(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgOrder)
		data[k] = tmp.OrderId
	}
	Ret.Data = data
}

func  ResolverOrderMain(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgOrderMain)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func  ResolverOrgBankCard(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgOrgBankCard)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func  ResolverOrgMoney(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgOrgMoney)
		data[k] = tmp.OrgId
	}
	Ret.Data = data
}

func  ResolverOrgMoneyBill(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgOrgMoneyBill)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func  ResolverOrgMoneyRecharge(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgOrgMoneyRecharge)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func  ResolverOrgMoneyWithdraw(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgOrgMoneyWithdraw)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func  ResolverOrganization(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgOrganization)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func  ResolverOrganizationActivity(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgOrganizationActivity)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func  ResolverOrganizationCar(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgOrganizationCar)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func  ResolverOrganizationCheck(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgOrganizationCheck)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func  ResolverOrganizationCheckInfo(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgOrganizationCheckInfo)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func  ResolverOrganizationConf(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgOrganizationConf)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func  ResolverOrganizationConfVal(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgOrganizationConfVal)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func  ResolverOrganizationCopy(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgOrganizationCopy)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func  ResolverOrganizationCopyCopy(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgOrganizationCopyCopy)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func  ResolverOrganizationFocus(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgOrganizationFocus)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func  ResolverOrganizationLog(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgOrganizationLog)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func  ResolverOrganizationOrder(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgOrganizationOrder)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func  ResolverOrganizationOrderpointsRecord(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgOrganizationOrderpointsRecord)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func  ResolverOrganizationPrivilege(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgOrganizationPrivilege)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func  ResolverOrganizationPrivilegeShop(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgOrganizationPrivilegeShop)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func  ResolverOrganizationPrivilegeStore(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgOrganizationPrivilegeStore)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func  ResolverOrganizationTime(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgOrganizationTime)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func  ResolverOrganizationType(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgOrganizationType)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func  ResolverPartner(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgPartner)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func  ResolverPartnerBlacklist(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgPartnerBlacklist)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func  ResolverPartnerInvite(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgPartnerInvite)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func  ResolverPartnerLog(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgPartnerLog)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func  ResolverPartnerPeriod(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgPartnerPeriod)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func  ResolverPartnerQrcode(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgPartnerQrcode)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func  ResolverPartnerRegister(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgPartnerRegister)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func  ResolverPartnerResults(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgPartnerResults)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func  ResolverPartnerTask(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgPartnerTask)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func  ResolverPayAlipayLog(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgPayAlipayLog)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func  ResolverPayWxLog(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgPayWxLog)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func  ResolverPhone(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgPhone)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func  ResolverPhoto(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgPhoto)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func  ResolverPointsData(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgPointsData)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func  ResolverPointsExceed(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgPointsExceed)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func  ResolverPointsExpRecords(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgPointsExpRecords)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func  ResolverPointsExpRules(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgPointsExpRules)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func  ResolverPointsGoods(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgPointsGoods)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func  ResolverPointsGoodsPay(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgPointsGoodsPay)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func  ResolverPointsLog(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgPointsLog)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func  ResolverPointsProduct(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgPointsProduct)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func  ResolverPointsProductRecords(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgPointsProductRecords)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func  ResolverPointsRanking(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgPointsRanking)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func  ResolverPointsSaleRules(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgPointsSaleRules)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func  ResolverPraise(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgPraise)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func  ResolverProjectExamine(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgProjectExamine)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func  ResolverPunch(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgPunch)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func  ResolverPurchase(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgPurchase)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func  ResolverPurchaseFactory(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgPurchaseFactory)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func  ResolverPurchaseFail(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgPurchaseFail)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func  ResolverPurchaseImg(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgPurchaseImg)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func  ResolverPurchaseRecords(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgPurchaseRecords)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func  ResolverPurchaseSms(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgPurchaseSms)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func  ResolverPurchaseUser(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgPurchaseUser)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func  ResolverPushSmsConfig(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgPushSmsConfig)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func  ResolverQuestionnaire(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgQuestionnaire)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func  ResolverQuestionnaireAnswer(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgQuestionnaireAnswer)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func  ResolverQuestionnaireMain(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgQuestionnaireMain)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func  ResolverQuestionnaireUser(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgQuestionnaireUser)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func  ResolverRabbitMsg(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgRabbitMsg)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func  ResolverRecommend(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgRecommend)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func  ResolverRecommendAuth(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgRecommendAuth)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func  ResolverRecommendContent(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgRecommendContent)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func  ResolverRecommendDetail201905(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgRecommendDetail201905)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func  ResolverRecommendDetail201906(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgRecommendDetail201906)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func  ResolverRecommendDetail201907(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgRecommendDetail201907)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func  ResolverRecommendDetail201908(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgRecommendDetail201908)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func  ResolverRecommendDetail201909(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgRecommendDetail201909)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func  ResolverRecommendDetail201910(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgRecommendDetail201910)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func  ResolverRecommendDetail201911(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgRecommendDetail201911)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func  ResolverRecommendDetail201912(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgRecommendDetail201912)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func  ResolverRecommendOrg(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgRecommendOrg)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func  ResolverRecommendPosition(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgRecommendPosition)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func  ResolverRecommendPositionConfig(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgRecommendPositionConfig)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func  ResolverRecommendRecord(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgRecommendRecord)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func  ResolverRecommendTime(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgRecommendTime)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func  ResolverRecommendTimeExtend(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgRecommendTimeExtend)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func  ResolverRecommendTongjiDay(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgRecommendTongjiDay)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func  ResolverRecommendTongjiMonth(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgRecommendTongjiMonth)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func  ResolverRedpack(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgRedpack)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func  ResolverRedpackArea(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgRedpackArea)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func  ResolverRedpackCheckRecord(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgRedpackCheckRecord)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func  ResolverRedpackContent(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgRedpackContent)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func  ResolverRedpackMsgRecord(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgRedpackMsgRecord)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func  ResolverRedpackOrg(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgRedpackOrg)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func  ResolverRedpackPosition(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgRedpackPosition)
		data[k] = tmp.Uid
	}
	Ret.Data = data
}

func  ResolverRedpackPrize(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgRedpackPrize)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func  ResolverRedpackRecord(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgRedpackRecord)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func  ResolverRedpackRecordTemp(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgRedpackRecordTemp)
		data[k] = tmp.Rid
	}
	Ret.Data = data
}

func  ResolverRedpackTime(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgRedpackTime)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func  ResolverRedpackUserCard(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgRedpackUserCard)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func  ResolverRedpackUserCardRecord(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgRedpackUserCardRecord)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func  ResolverRegion(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgRegion)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func  ResolverReport(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgReport)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func  ResolverReportImg(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgReportImg)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func  ResolverReportRecord(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgReportRecord)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func  ResolverRide(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgRide)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func  ResolverRideCheck(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgRideCheck)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func  ResolverRideCheckNotes(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgRideCheckNotes)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func  ResolverRideLine(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgRideLine)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func  ResolverRideNode(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgRideNode)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func  ResolverRoadBook(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgRoadBook)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func  ResolverRoadBookLine(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgRoadBookLine)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func  ResolverRoadBookLineContent(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgRoadBookLineContent)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func  ResolverRoadBookMeet(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgRoadBookMeet)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func  ResolverRoadBookRecommend(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgRoadBookRecommend)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func  ResolverRoadbookArouse(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgRoadbookArouse)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func  ResolverRoadbookArouseAbnormal(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgRoadbookArouseAbnormal)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func  ResolverRoadbookArouseAccount(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgRoadbookArouseAccount)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func  ResolverRoadbookArouseAttend(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgRoadbookArouseAttend)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func  ResolverRoadbookArouseAttendCensus(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgRoadbookArouseAttendCensus)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func  ResolverRoadbookArouseFlow(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgRoadbookArouseFlow)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func  ResolverRoadbookArouseLog(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgRoadbookArouseLog)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func  ResolverRoadbookArouseRealAccount(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgRoadbookArouseRealAccount)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func  ResolverRoadbookArouseRelated(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgRoadbookArouseRelated)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func  ResolverRoadbookLine(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgRoadbookLine)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func  ResolverRoadbookMerchant(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgRoadbookMerchant)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func  ResolverRoadbookMerchantApply(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgRoadbookMerchantApply)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func  ResolverRoadbookMerchantInvite(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgRoadbookMerchantInvite)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func  ResolverRoadbookNodePunchCensus(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgRoadbookNodePunchCensus)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func  ResolverRouletteLottery(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgRouletteLottery)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func  ResolverRouletteLotteryAttend(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgRouletteLotteryAttend)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func  ResolverRouletteLotteryAward(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgRouletteLotteryAward)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func  ResolverSaleinfoOrg(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgSaleinfoOrg)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func  ResolverSalerStar(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgSalerStar)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func  ResolverSearchLog(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgSearchLog)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func  ResolverSeries(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgSeries)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func  ResolverServiceCity(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgServiceCity)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func  ResolverSmsApiLog(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgSmsApiLog)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func  ResolverSmsTemplate(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgSmsTemplate)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func  ResolverSmsTemplateLog(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgSmsTemplateLog)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func  ResolverSmsTotal(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgSmsTotal)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func  ResolverSmsVerifyCode(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgSmsVerifyCode)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func  ResolverSpecial(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgSpecial)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func  ResolverSpmAppDownloadLog(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgSpmAppDownloadLog)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func  ResolverSpmBaidurefererLog(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgSpmBaidurefererLog)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func  ResolverSpmCate(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgSpmCate)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func  ResolverSpmDesignation(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgSpmDesignation)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func  ResolverSpmKeywords(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgSpmKeywords)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func  ResolverSpmLog(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgSpmLog)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func  ResolverSpmLogCopy(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgSpmLogCopy)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func  ResolverSpmParameter(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgSpmParameter)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func  ResolverSpmPlace(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgSpmPlace)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func  ResolverSpmPromoter(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgSpmPromoter)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func  ResolverStaff(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgStaff)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func  ResolverStatistics(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgStatistics)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func  ResolverSupertopicApply(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgSupertopicApply)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func  ResolverTagCar(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgTagCar)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func  ResolverTagCartSpecific(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgTagCartSpecific)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func  ResolverTags(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgTags)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func  ResolverTagsCategory(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgTagsCategory)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func  ResolverTagscategory(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgTagscategory)
		data[k] = tmp.Id
	}
	Ret.Data = data
}


func ResolverThreadActivity(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgThreadActivity)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func  ResolverThreadActivityApply(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgThreadActivityApply)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func  ResolverThreadActivityCar(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgThreadActivityCar)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func  ResolverThreadActivityMaster(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgThreadActivityMaster)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func  ResolverThreadActivityScene(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgThreadActivityScene)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func  ResolverThreadActivityType(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgThreadActivityType)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func  ResolverThreadActivityUser(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgThreadActivityUser)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func  ResolverThreadAsk(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgThreadAsk)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func  ResolverThreadAskAnswer(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgThreadAskAnswer)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func  ResolverThreadAskAnswerFollow(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgThreadAskAnswerFollow)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func  ResolverThreadCensusRanking(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgThreadCensusRanking)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func  ResolverThreadCycle(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgThreadCycle)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func  ResolverThreadDeleteReason(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgThreadDeleteReason)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func  ResolverThreadEvaluate(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgThreadEvaluate)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func  ResolverThreadEvaluateCar(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgThreadEvaluateCar)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func  ResolverThreadGoods(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgThreadGoods)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func  ResolverThreadGoodsCategory(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgThreadGoodsCategory)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func  ResolverThreadIntegralLog(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgThreadIntegralLog)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func  ResolverThreadMastersay(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgThreadMastersay)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func  ResolverThreadMastersayCar(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgThreadMastersayCar)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func  ResolverThreadMood(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgThreadMood)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func  ResolverThreadPlay(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgThreadPlay)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func  ResolverThreadRoadbookAnnex(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgThreadRoadbookAnnex)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func  ResolverThreadRoadbookLineBeen(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgThreadRoadbookLineBeen)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func  ResolverThreadRoadbookLineRelated(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgThreadRoadbookLineRelated)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func  ResolverThreadRoadbookNoteUser(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgThreadRoadbookNoteUser)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func  ResolverThreadRoadbookPunchCensus(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgThreadRoadbookPunchCensus)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func  ResolverThreadShow(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgThreadShow)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func  ResolverThreadTopic(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgThreadTopic)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func  ResolverThreadTopicAccess(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgThreadTopicAccess)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func  ResolverThreadTopicApply(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgThreadTopicApply)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func  ResolverThreadTopicFans(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgThreadTopicFans)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func  ResolverThreadVisible(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgThreadVisible)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func  ResolverThreadWordColumn(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgThreadWordColumn)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func  ResolverThreadWords(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgThreadWords)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func  ResolverTicket(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgTicket)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func  ResolverTicketApiLog(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgTicketApiLog)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func  ResolverTicketApplyUser(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgTicketApplyUser)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func  ResolverTicketCheckRecord(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgTicketCheckRecord)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func  ResolverTicketCheckoutAuth(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgTicketCheckoutAuth)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func  ResolverTicketFree(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgTicketFree)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func  ResolverTicketFreeHandout(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgTicketFreeHandout)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func  ResolverTicketNumsRecord(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgTicketNumsRecord)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func  ResolverTicketPaperRecord(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgTicketPaperRecord)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func  ResolverTicketQd(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgTicketQd)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func  ResolverTicketQdSale(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgTicketQdSale)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func  ResolverTicketRedpack(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgTicketRedpack)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func  ResolverTicketRedpackLog(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgTicketRedpackLog)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func  ResolverTicketRefund(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgTicketRefund)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func  ResolverTicketRefundLog(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgTicketRefundLog)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func  ResolverTicketSendRecord(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgTicketSendRecord)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func  ResolverTicketThirdParty(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgTicketThirdParty)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func  ResolverTicketUser(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgTicketUser)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func  ResolverTmpData(multiModels []interface{}, length int )  {
	data := make([]int, length)
	//for k, model:= range multiModels  {
	//	tmp := (model).(models.TgTmpData)
	//	data[k] = tmp.Id
	//}
	Ret.Data = data
}

func  ResolverTopicCategory(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgTopicCategory)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func  ResolverTopicExpLog(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgTopicExpLog)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func  ResolverTopicLog(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgTopicLog)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func  ResolverTopicScompere(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgTopicScompere)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func  ResolverTopicSignin(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgTopicSignin)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func  ResolverTrafficStatistics(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgTrafficStatistics)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func  ResolverTrafficStatisticsLog(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgTrafficStatisticsLog)
		data[k] = tmp.LogId
	}
	Ret.Data = data
}

func  ResolverTrafficStatisticsLogGroup(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgTrafficStatisticsLogGroup)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func  ResolverTwitter(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgTwitter)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func  ResolverTwitterAlbum(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgTwitterAlbum)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func  ResolverTwitterSource(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgTwitterSource)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func  ResolverTwitterTopicAccess(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgTwitterTopicAccess)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func  ResolverTwitterVisibility(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgTwitterVisibility)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func  ResolverUploadApk(multiModels []interface{}, length int )  {
	data := make([]int, length)
	//for k, model:= range multiModels  {
	//	tmp := (model).(models.TgUploadApk)
	//	//data[k] = tmp.Id
	//}
	Ret.Data = data
}

func  ResolverUserBankCard(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgUserBankCard)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func  ResolverUserBehaviour(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgUserBehaviour)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func  ResolverUserCar(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgUserCar)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func  ResolverUserFollow(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgUserFollow)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func  ResolverUserPoints(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgUserPoints)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func  ResolverUserRecommend(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgUserRecommend)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func  ResolverUserSetting(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgUserSetting)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func  ResolverVideo(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgVideo)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func  ResolverVideoAttach(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgVideoAttach)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func  ResolverVideoAttachReceive(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgVideoAttachReceive)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func  ResolverVideoLog(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgVideoLog)
		data[k] = tmp.LogId
	}
	Ret.Data = data
}

func  ResolverVideoNotice(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgVideoNotice)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func  ResolverVideoRecommendTable(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgVideoRecommendTable)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func  ResolverVideoVisible(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgVideoVisible)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func  ResolverVirtualNickname(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgVirtualNickname)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func  ResolverVirtualRecord(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgVirtualRecord)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func  ResolverWannaBuy(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgWannaBuy)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func  ResolverWeatherCitys(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgWeatherCitys)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func  ResolverWeatherToday(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgWeatherToday)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func  ResolverWeixinConfig(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgWeixinConfig)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func  ResolverWeixinCustomSend(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgWeixinCustomSend)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func  ResolverWeixinKerword(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgWeixinKerword)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func  ResolverWeixinMaterial(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgWeixinMaterial)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func  ResolverWeixinMenu(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgWeixinMenu)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func  ResolverWeixinPush(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgWeixinPush)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func  ResolverWeixinPushAuto(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgWeixinPushAuto)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func  ResolverWeixinPushConfig(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgWeixinPushConfig)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func  ResolverWeixinPushInteractive(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgWeixinPushInteractive)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func  ResolverWeixinPushMaterial(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgWeixinPushMaterial)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func  ResolverWeixinPushRecord(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgWeixinPushRecord)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func  ResolverWeixinPushTag(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgWeixinPushTag)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func  ResolverWeixinQrcode(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgWeixinQrcode)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func  ResolverWeixinQrcodeScanLog(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgWeixinQrcodeScanLog)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func  ResolverWeixinTemplate(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgWeixinTemplate)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func  ResolverXc1Jurisdiction(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgXc1Jurisdiction)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func  ResolverXc1ManageObject(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgXc1ManageObject)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func  ResolverXc1Org2scene(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgXc1Org2scene)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func  ResolverXc1Worker2guest(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgXc1Worker2guest)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func  ResolverXcAppMessage(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgXcAppMessage)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func  ResolverXcAppMessageCount(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgXcAppMessageCount)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func  ResolverXcApplyItem(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgXcApplyItem)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func  ResolverXcApplyType(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgXcApplyType)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func  ResolverXcArea(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgXcArea)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func  ResolverXcAuthMapping(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgXcAuthMapping)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func  ResolverXcAuthorise(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgXcAuthorise)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func  ResolverXcCheckCourse(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgXcCheckCourse)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func  ResolverXcCheckTable(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgXcCheckTable)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func  ResolverXcChecker2object(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgXcChecker2object)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func  ResolverXcFriends(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgXcFriends)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func  ResolverXcGuest2object(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgXcGuest2object)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func  ResolverXcJobType(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgXcJobType)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func  ResolverXcMessage(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgXcMessage)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func  ResolverXcMessageAddtask(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgXcMessageAddtask)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func  ResolverXcMessageApply(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgXcMessageApply)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func  ResolverXcMessageContent(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgXcMessageContent)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func  ResolverXcMessageRectification(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgXcMessageRectification)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func  ResolverXcMessageSetting(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgXcMessageSetting)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func  ResolverXcMessageUnqualifiedcheck(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgXcMessageUnqualifiedcheck)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func  ResolverXcObject(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgXcObject)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func  ResolverXcObject2dispose(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgXcObject2dispose)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func  ResolverXcObjectOrganization(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgXcObjectOrganization)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func  ResolverXcObjectType(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgXcObjectType)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func  ResolverXcOffsetTriggerTime(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgXcOffsetTriggerTime)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func  ResolverXcOrder(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgXcOrder)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func  ResolverXcOrderProgress(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgXcOrderProgress)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func  ResolverXcOrganization(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgXcOrganization)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func  ResolverXcPayLog(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgXcPayLog)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func  ResolverXcPermission(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgXcPermission)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func  ResolverXcScene(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgXcScene)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func  ResolverXcSceneReview(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgXcSceneReview)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func  ResolverXcStaff2object(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgXcStaff2object)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func  ResolverXcTask(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgXcTask)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func  ResolverXcTaskCopy(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgXcTaskCopy)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func  ResolverXcTaskDis(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgXcTaskDis)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func  ResolverXcTaskDisimg(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgXcTaskDisimg)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func  ResolverXcTaskLog(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgXcTaskLog)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func  ResolverXcTaskPackage(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgXcTaskPackage)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func  ResolverXcTopUpLog(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgXcTopUpLog)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func  ResolverXcUser2scene(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgXcUser2scene)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func  ResolverXcUsercoordinate(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgXcUsercoordinate)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func  ResolverXcWorker(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgXcWorker)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func  ResolverXcWorker2area(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgXcWorker2area)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func  ResolverXcWorkerYanghuanBack(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgXcWorkerYanghuanBack)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func  ResolverYaoyiyao(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgYaoyiyao)
		data[k] = tmp.YyyId
	}
	Ret.Data = data
}

func  ResolverYaoyiyaoAuth(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgYaoyiyaoAuth)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func  ResolverYaoyiyaoAwards(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgYaoyiyaoAwards)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func  ResolverYaoyiyaoAwardsCopy(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgYaoyiyaoAwardsCopy)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func  ResolverYaoyiyaoCacheawards(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgYaoyiyaoCacheawards)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func  ResolverYaoyiyaoEndLog(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgYaoyiyaoEndLog)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func  ResolverYaoyiyaoNotWin(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgYaoyiyaoNotWin)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func  ResolverYaoyiyaoPassLog(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgYaoyiyaoPassLog)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func  ResolverYaoyiyaoRank(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgYaoyiyaoRank)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func  ResolverYaoyiyaoRotary(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgYaoyiyaoRotary)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func  ResolverYaoyiyaoSignin(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgYaoyiyaoSignin)
		data[k] = tmp.Id
	}
	Ret.Data = data
}

func  ResolverYaoyiyaoStatus(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(models.TgYaoyiyaoStatus)
		data[k] = tmp.Id
	}
	Ret.Data = data
}
