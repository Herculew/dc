//table 映射器
package service

import (
	"dc/models"
	"github.com/mitchellh/mapstructure"
)
	//table - model的映射

var	TableBindToModel = map[string]func(args Args)interface{}{
		tablePrefix+"thread": BindThread,
		tablePrefix+"thread_roadbook": BindThreadRoadbook,
		tablePrefix+"2handcar_apply": Bind2handcarApply,
		tablePrefix+"2handcar_booth":Bind2handcarBooth,
		tablePrefix+"2handcar_complain":Bind2handcarComplain,
		tablePrefix+"2handcar_config":Bind2handcarConfig,
		tablePrefix+"2handcar_debit_list":Bind2handcarDebitList,
		tablePrefix+"2handcar_intention":Bind2handcarIntention,
		tablePrefix+"2handcar_list":Bind2handcarList,
		tablePrefix+"2handcar_order":Bind2handcarOrder,
		tablePrefix+"2handcar_park":Bind2handcarPark,
		tablePrefix+"2handcar_park_time":Bind2handcarParkTime,
		tablePrefix+"2handcar_park_type":Bind2handcarParkType,
		tablePrefix+"2handcar_plan":Bind2handcarPlan,
		tablePrefix+"2handcar_time":Bind2handcarTime,
		tablePrefix+"2handcar_wx_fail":Bind2handcarWxFail,
		tablePrefix+"account_manage":BindAccountManage,
		tablePrefix+"action_log":BindActionLog,
		tablePrefix+"activity_area":BindActivityArea,
		tablePrefix+"activity_forward":BindActivityForward,
		tablePrefix+"activity_forward_content":BindActivityForwardContent,
		tablePrefix+"activity_forward_prize":BindActivityForwardPrize,
		tablePrefix+"activity_forward_record":BindActivityForwardRecord,
		tablePrefix+"activity_forward_record_details":BindActivityForwardRecordDetails,
		tablePrefix+"activity_line":BindActivityLine,
		tablePrefix+"activity_line_node_related":BindActivityLineNodeRelated,
		tablePrefix+"activity_lottery_attendance":BindActivityLotteryAttendance,
		tablePrefix+"activity_lottery_auth":BindActivityLotteryAuth,
		tablePrefix+"activity_lottery_code":BindActivityLotteryCode,
		tablePrefix+"activity_lottery_prize":BindActivityLotteryPrize,
		tablePrefix+"activity_lottery_rotary":BindActivityLotteryRotary,
		tablePrefix+"activity_lottery_sms":BindActivityLotterySms,
		tablePrefix+"activity_lottery_status":BindActivityLotteryStatus,
		tablePrefix+"activity_table":BindActivityTable,
		tablePrefix+"activity_tags":BindActivityTags,
		tablePrefix+"activity_turntable":BindActivityTurntable,
		tablePrefix+"activity_turntable_content":BindActivityTurntableContent,
		tablePrefix+"activity_turntable_prize":BindActivityTurntablePrize,
		tablePrefix+"activity_turntable_qrcodetag":BindActivityTurntableQrcodetag,
		tablePrefix+"activity_turntable_record":BindActivityTurntableRecord,
		tablePrefix+"activity_turntable_result":BindActivityTurntableResult,
		tablePrefix+"activity_turntable_saddress":BindActivityTurntableSaddress,
		tablePrefix+"activity_turntable_share":BindActivityTurntableShare,
		tablePrefix+"activity_turntable_user":BindActivityTurntableUser,
		tablePrefix+"activity_vote_refer_opt":BindActivityVoteReferOpt,
		tablePrefix+"activity_vote_rules":BindActivityVoteRules,
		tablePrefix+"activity_vote_user_limit":BindActivityVoteUserLimit,
		tablePrefix+"activity_vote_user_table":BindActivityVoteUserTable,
		tablePrefix+"activity_vote_user_ticketcert":BindActivityVoteUserTicketcert,
		tablePrefix+"admin_auth_project_xc":BindAdminAuthProjectXc,
		tablePrefix+"admin_auth_project_xc_copy":BindAdminAuthProjectXcCopy,
		tablePrefix+"admin_auth_user":BindAdminAuthUser,
		tablePrefix+"admin_auth_user_copy":BindAdminAuthUserCopy,
		tablePrefix+"admin_menu":BindAdminMenu,
		tablePrefix+"admin_report":BindAdminReport,
		tablePrefix+"admin_role":BindAdminRole,
		tablePrefix+"admin_user":BindAdminUser,
		tablePrefix+"admin_user2area":BindAdminUser2area,
		tablePrefix+"admin_user2area_org":BindAdminUser2areaOrg,
		tablePrefix+"admin_user2org":BindAdminUser2org,
		tablePrefix+"admin_user2org_copy":BindAdminUser2orgCopy,
		tablePrefix+"admin_user2org_setting":BindAdminUser2orgSetting,
		tablePrefix+"admin_user_alltags":BindAdminUserAlltags,
		tablePrefix+"admin_user_auth":BindAdminUserAuth,
		tablePrefix+"admin_user_auth_copy":BindAdminUserAuthCopy,
		tablePrefix+"admin_user_avatar":BindAdminUserAvatar,
		tablePrefix+"admin_user_copy1":BindAdminUserCopy1,
		tablePrefix+"admin_user_department":BindAdminUserDepartment,
		tablePrefix+"admin_user_doyen":BindAdminUserDoyen,
		tablePrefix+"admin_user_favor":BindAdminUserFavor,
		tablePrefix+"admin_user_impression":BindAdminUserImpression,
		tablePrefix+"admin_user_lock":BindAdminUserLock,
		tablePrefix+"admin_user_login_fail":BindAdminUserLoginFail,
		tablePrefix+"admin_user_money":BindAdminUserMoney,
		tablePrefix+"admin_user_money_bill":BindAdminUserMoneyBill,
		tablePrefix+"admin_user_money_recharge":BindAdminUserMoneyRecharge,
		tablePrefix+"admin_user_money_withdraw":BindAdminUserMoneyWithdraw,
		tablePrefix+"admin_user_org":BindAdminUserOrg,
		tablePrefix+"admin_user_position":BindAdminUserPosition,
		tablePrefix+"admin_user_punish":BindAdminUserPunish,
		tablePrefix+"admin_user_real_name_auth":BindAdminUserRealNameAuth,
		tablePrefix+"admin_user_role":BindAdminUserRole,
		tablePrefix+"admin_user_tags":BindAdminUserTags,
		tablePrefix+"admin_user_type":BindAdminUserType,
		tablePrefix+"advertisement_admin_log":BindAdvertisementAdminLog,
		tablePrefix+"advertisement_content":BindAdvertisementContent,
		tablePrefix+"advertisement_org_strategy":BindAdvertisementOrgStrategy,
		tablePrefix+"advertisement_position":BindAdvertisementPosition,
		tablePrefix+"agreement":BindAgreement,
		tablePrefix+"agreement_place":BindAgreementPlace,
		tablePrefix+"album":BindAlbum,
		tablePrefix+"alltags":BindAlltags,
		tablePrefix+"alltags_article":BindAlltagsArticle,
		tablePrefix+"alltags_category":BindAlltagsCategory,
		tablePrefix+"alltags_category_activity":BindAlltagsCategoryActivity,
		tablePrefix+"area":BindArea,
		tablePrefix+"area_bak":BindAreaBak,
		tablePrefix+"area_hot":BindAreaHot,
		tablePrefix+"article":BindArticle,
		tablePrefix+"article_category":BindArticleCategory,
		tablePrefix+"article_data":BindArticleData,
		tablePrefix+"article_tag":BindArticleTag,
		tablePrefix+"ask_cate":BindAskCate,
		tablePrefix+"ask_expert":BindAskExpert,
		tablePrefix+"async_task":BindAsyncTask,
		tablePrefix+"attachment":BindAttachment,
		tablePrefix+"auth_action":BindAuthAction,
		tablePrefix+"auth_role":BindAuthRole,
		tablePrefix+"auth_role_action":BindAuthRoleAction,
		tablePrefix+"auth_role_source":BindAuthRoleSource,
		tablePrefix+"auth_user_role":BindAuthUserRole,
		tablePrefix+"baidu_config":BindBaiduConfig,
		tablePrefix+"baidu_notice":BindBaiduNotice,
		tablePrefix+"baidu_share":BindBaiduShare,
		tablePrefix+"benefit":BindBenefit,
		tablePrefix+"benefit_access":BindBenefitAccess,
		tablePrefix+"benefit_store":BindBenefitStore,
		tablePrefix+"bid_keyword":BindBidKeyword,
		tablePrefix+"bid_keyword_copy":BindBidKeywordCopy,
		tablePrefix+"brand":BindBrand,
		tablePrefix+"browse_history":BindBrowseHistory,
		tablePrefix+"call_audience":BindCallAudience,
		tablePrefix+"call_audience_access":BindCallAudienceAccess,
		tablePrefix+"call_audience_conf":BindCallAudienceConf,
		tablePrefix+"call_audience_conf_val":BindCallAudienceConfVal,
		tablePrefix+"call_audience_review":BindCallAudienceReview,
		tablePrefix+"call_audience_review_log":BindCallAudienceReviewLog,
		tablePrefix+"call_audience_smstask":BindCallAudienceSmstask,
		tablePrefix+"call_audience_smstask_content":BindCallAudienceSmstaskContent,
		tablePrefix+"call_audience_smstask_record":BindCallAudienceSmstaskRecord,
		tablePrefix+"call_audience_user":BindCallAudienceUser,
		tablePrefix+"call_audience_visit":BindCallAudienceVisit,
		tablePrefix+"call_datatype":BindCallDatatype,
		tablePrefix+"call_draw":BindCallDraw,
		tablePrefix+"call_event_auth":BindCallEventAuth,
		tablePrefix+"call_privacy":BindCallPrivacy,
		tablePrefix+"call_questionnaire":BindCallQuestionnaire,
		tablePrefix+"call_questionnaire_conf":BindCallQuestionnaireConf,
		tablePrefix+"call_questionnaire_user":BindCallQuestionnaireUser,
		tablePrefix+"call_questionnaire_val":BindCallQuestionnaireVal,
		tablePrefix+"call_record":BindCallRecord,
		tablePrefix+"call_user_auth":BindCallUserAuth,
		tablePrefix+"car":BindCar,
		tablePrefix+"car2conf":BindCar2conf,
		tablePrefix+"car_color":BindCarColor,
		tablePrefix+"car_conf":BindCarConf,
		tablePrefix+"car_conf_val":BindCarConfVal,
		tablePrefix+"car_copy":BindCarCopy,
		tablePrefix+"car_custom_made":BindCarCustomMade,
		tablePrefix+"car_find":BindCarFind,
		tablePrefix+"car_love":BindCarLove,
		tablePrefix+"car_order":BindCarOrder,
		tablePrefix+"car_order_config":BindCarOrderConfig,
		tablePrefix+"car_order_list":BindCarOrderList,
		tablePrefix+"car_order_msg":BindCarOrderMsg,
		tablePrefix+"car_order_offer":BindCarOrderOffer,
		tablePrefix+"car_order_offer_count":BindCarOrderOfferCount,
		tablePrefix+"car_order_push_list":BindCarOrderPushList,
		tablePrefix+"car_order_tags":BindCarOrderTags,
		tablePrefix+"car_pic":BindCarPic,
		tablePrefix+"car_search_log":BindCarSearchLog,
		tablePrefix+"car_series":BindCarSeries,
		tablePrefix+"car_series_color":BindCarSeriesColor,
		tablePrefix+"carorder":BindCarorder,
		tablePrefix+"carorder_config":BindCarorderConfig,
		tablePrefix+"carorder_list":BindCarorderList,
		tablePrefix+"carorder_lovecar":BindCarorderLovecar,
		tablePrefix+"carorder_offer":BindCarorderOffer,
		tablePrefix+"carorder_offer_record":BindCarorderOfferRecord,
		tablePrefix+"carorder_push_list":BindCarorderPushList,
		tablePrefix+"carorder_record":BindCarorderRecord,
		tablePrefix+"carorder_tactics":BindCarorderTactics,
		tablePrefix+"carorder_tactics_config":BindCarorderTacticsConfig,
		tablePrefix+"carorder_tactics_gift":BindCarorderTacticsGift,
		tablePrefix+"carorder_tactics_record":BindCarorderTacticsRecord,
		tablePrefix+"carousel":BindCarousel,
		tablePrefix+"carousel_item":BindCarouselItem,
		tablePrefix+"carousel_object":BindCarouselObject,
		tablePrefix+"carousel_org_plan":BindCarouselOrgPlan,
		tablePrefix+"carousel_org_right":BindCarouselOrgRight,
		tablePrefix+"carousel_originality":BindCarouselOriginality,
		tablePrefix+"carousel_scheduling":BindCarouselScheduling,
		tablePrefix+"carousel_scheduling_authorize":BindCarouselSchedulingAuthorize,
		tablePrefix+"carousel_strategy":BindCarouselStrategy,
		tablePrefix+"carousel_subject_item":BindCarouselSubjectItem,
		tablePrefix+"cert":BindCert,
		tablePrefix+"cert_apply":BindCertApply,
		tablePrefix+"cert_content":BindCertContent,
		tablePrefix+"cert_handout":BindCertHandout,
		tablePrefix+"cert_nums_record":BindCertNumsRecord,
		tablePrefix+"cert_user":BindCertUser,
		tablePrefix+"cert_user_copy":BindCertUserCopy,
		tablePrefix+"cert_user_record":BindCertUserRecord,
		tablePrefix+"cert_user_record_act":BindCertUserRecordAct,
		tablePrefix+"chicken_soup":BindChickenSoup,
		tablePrefix+"collect":BindCollect,
		tablePrefix+"collect_mapping":BindCollectMapping,
		tablePrefix+"comment":BindComment,
		tablePrefix+"comment_log":BindCommentLog,
		tablePrefix+"config":BindConfig,
		tablePrefix+"country":BindCountry,
		tablePrefix+"coupon":BindCoupon,
		tablePrefix+"coupon_car":BindCouponCar,
		tablePrefix+"coupon_category":BindCouponCategory,
		tablePrefix+"coupon_check":BindCouponCheck,
		tablePrefix+"coupon_code":BindCouponCode,
		tablePrefix+"coupon_dealers":BindCouponDealers,
		tablePrefix+"coupon_distribute":BindCouponDistribute,
		tablePrefix+"coupon_distribute_list":BindCouponDistributeList,
		tablePrefix+"coupon_distribute_list_reissue":BindCouponDistributeListReissue,
		tablePrefix+"coupon_distribute_reissue":BindCouponDistributeReissue,
		tablePrefix+"coupon_issue":BindCouponIssue,
		tablePrefix+"coupon_order":BindCouponOrder,
		tablePrefix+"coupon_push":BindCouponPush,
		tablePrefix+"coupon_push_log":BindCouponPushLog,
		tablePrefix+"coupon_refund":BindCouponRefund,
		tablePrefix+"coupon_replace":BindCouponReplace,
		tablePrefix+"coupon_replace_distribute":BindCouponReplaceDistribute,
		tablePrefix+"coupon_replace_reissue":BindCouponReplaceReissue,
		tablePrefix+"coupon_share_log":BindCouponShareLog,
		tablePrefix+"coupon_ticket":BindCouponTicket,
		tablePrefix+"coupon_warehouse":BindCouponWarehouse,
		tablePrefix+"coupon_warehouse_code":BindCouponWarehouseCode,
		tablePrefix+"crawler_data":BindCrawlerData,
		tablePrefix+"danger":BindDanger,
		tablePrefix+"dshow_banner":BindDshowBanner,
		tablePrefix+"dshow_base":BindDshowBase,
		tablePrefix+"dshow_gift":BindDshowGift,
		tablePrefix+"dshow_share":BindDshowShare,
		tablePrefix+"dshow_text":BindDshowText,
		tablePrefix+"easemob_messages":BindEasemobMessages,
		tablePrefix+"easemob_messages_num":BindEasemobMessagesNum,
		tablePrefix+"easemob_messages_payload":BindEasemobMessagesPayload,
		tablePrefix+"emoji":BindEmoji,
		tablePrefix+"emoji_group":BindEmojiGroup,
		tablePrefix+"emoji_user_access":BindEmojiUserAccess,
		tablePrefix+"entrance_way":BindEntranceWay,
		tablePrefix+"event":BindEvent,
		tablePrefix+"event_activity":BindEventActivity,
		tablePrefix+"event_activity_data":BindEventActivityData,
		tablePrefix+"event_ad":BindEventAd,
		tablePrefix+"event_bdconfig":BindEventBdconfig,
		tablePrefix+"event_booth":BindEventBooth,
		tablePrefix+"event_brand":BindEventBrand,
		tablePrefix+"event_brand_merge":BindEventBrandMerge,
		tablePrefix+"event_business":BindEventBusiness,
		tablePrefix+"event_desc_img":BindEventDescImg,
		tablePrefix+"event_examine":BindEventExamine,
		tablePrefix+"event_guest":BindEventGuest,
		tablePrefix+"event_hall":BindEventHall,
		tablePrefix+"event_img_config":BindEventImgConfig,
		tablePrefix+"event_img_unshow":BindEventImgUnshow,
		tablePrefix+"event_information":BindEventInformation,
		tablePrefix+"event_model":BindEventModel,
		tablePrefix+"event_navigation":BindEventNavigation,
		tablePrefix+"event_navigation_conf":BindEventNavigationConf,
		tablePrefix+"event_newcar":BindEventNewcar,
		tablePrefix+"event_schedule":BindEventSchedule,
		tablePrefix+"event_showcar":BindEventShowcar,
		tablePrefix+"event_staff":BindEventStaff,
		tablePrefix+"event_staff_auth":BindEventStaffAuth,
		tablePrefix+"event_staff_auth_area":BindEventStaffAuthArea,
		tablePrefix+"event_staff_auth_brand":BindEventStaffAuthBrand,
		tablePrefix+"event_staff_auth_user":BindEventStaffAuthUser,
		tablePrefix+"event_user":BindEventUser,
		tablePrefix+"fab":BindFab,
		tablePrefix+"factory":BindFactory,
		tablePrefix+"failed_jobs":BindFailedJobs,
		tablePrefix+"feedback":BindFeedback,
		tablePrefix+"feeds_ignore":BindFeedsIgnore,
		tablePrefix+"feeds_topic":BindFeedsTopic,
		tablePrefix+"feeds_user":BindFeedsUser,
		tablePrefix+"follow_group":BindFollowGroup,
		tablePrefix+"friend":BindFriend,
		tablePrefix+"friend_buyinfo":BindFriendBuyinfo,
		tablePrefix+"friend_group":BindFriendGroup,
		tablePrefix+"friend_group_access":BindFriendGroupAccess,
		tablePrefix+"friend_latent":BindFriendLatent,
		tablePrefix+"friend_manage_log":BindFriendManageLog,
		tablePrefix+"friend_org":BindFriendOrg,
		tablePrefix+"gallery":BindGallery,
		tablePrefix+"group":BindGroup,
		tablePrefix+"group_authentication":BindGroupAuthentication,
		tablePrefix+"group_blacklist":BindGroupBlacklist,
		tablePrefix+"group_category":BindGroupCategory,
		tablePrefix+"group_invitation":BindGroupInvitation,
		tablePrefix+"group_notification":BindGroupNotification,
		tablePrefix+"group_notification_read":BindGroupNotificationRead,
		tablePrefix+"group_qrcode":BindGroupQrcode,
		tablePrefix+"group_tag":BindGroupTag,
		tablePrefix+"group_tag_access":BindGroupTagAccess,
		tablePrefix+"group_user":BindGroupUser,
		tablePrefix+"group_user_kick_out":BindGroupUserKickOut,
		tablePrefix+"help_question":BindHelpQuestion,
		tablePrefix+"hot_search":BindHotSearch,
		tablePrefix+"hot_search_ext":BindHotSearchExt,
		tablePrefix+"img":BindImg,
		tablePrefix+"img_tag":BindImgTag,
		tablePrefix+"img_type":BindImgType,
		tablePrefix+"industry":BindIndustry,
		tablePrefix+"industry_user":BindIndustryUser,
		tablePrefix+"koubei":BindKoubei,
		tablePrefix+"koubei_album":BindKoubeiAlbum,
		tablePrefix+"koubei_spider":BindKoubeiSpider,
		tablePrefix+"latent_autoset":BindLatentAutoset,
		tablePrefix+"latent_dispense":BindLatentDispense,
		tablePrefix+"line":BindLine,
		tablePrefix+"line_node":BindLineNode,
		tablePrefix+"line_node_album":BindLineNodeAlbum,
		tablePrefix+"line_node_comment":BindLineNodeComment,
		tablePrefix+"line_node_comment_minutia":BindLineNodeCommentMinutia,
		tablePrefix+"line_node_comment_text":BindLineNodeCommentText,
		tablePrefix+"line_node_punch_census":BindLineNodePunchCensus,
		tablePrefix+"line_node_related":BindLineNodeRelated,
		tablePrefix+"live_message":BindLiveMessage,
		tablePrefix+"live_user":BindLiveUser,
		tablePrefix+"log_admin_login":BindLogAdminLogin,
		tablePrefix+"log_admin_work":BindLogAdminWork,
		tablePrefix+"lottery":BindLottery,
		tablePrefix+"lottery_auth":BindLotteryAuth,
		tablePrefix+"lottery_config":BindLotteryConfig,
		tablePrefix+"lottery_cron":BindLotteryCron,
		tablePrefix+"lottery_new":BindLotteryNew,
		tablePrefix+"lottery_new_auth":BindLotteryNewAuth,
		tablePrefix+"lottery_new_code_bill":BindLotteryNewCodeBill,
		tablePrefix+"lottery_new_config":BindLotteryNewConfig,
		tablePrefix+"lottery_new_cron":BindLotteryNewCron,
		tablePrefix+"lottery_new_lucky_code":BindLotteryNewLuckyCode,
		tablePrefix+"lottery_new_order":BindLotteryNewOrder,
		tablePrefix+"lottery_new_org_code":BindLotteryNewOrgCode,
		tablePrefix+"lottery_new_prize":BindLotteryNewPrize,
		tablePrefix+"lottery_new_record":BindLotteryNewRecord,
		tablePrefix+"lottery_new_set":BindLotteryNewSet,
		tablePrefix+"lottery_new_user":BindLotteryNewUser,
		tablePrefix+"lottery_new_user_code":BindLotteryNewUserCode,
		tablePrefix+"lottery_org_code":BindLotteryOrgCode,
		tablePrefix+"lottery_part":BindLotteryPart,
		tablePrefix+"lottery_prize":BindLotteryPrize,
		tablePrefix+"lottery_record":BindLotteryRecord,
		tablePrefix+"lottery_thread":BindLotteryThread,
		tablePrefix+"lottery_user":BindLotteryUser,
		tablePrefix+"lottery_user_code":BindLotteryUserCode,
		tablePrefix+"love_chicken":BindLoveChicken,
		tablePrefix+"lover_user":BindLoverUser,
		tablePrefix+"main_work":BindMainWork,
		tablePrefix+"management_type":BindManagementType,
		tablePrefix+"manual_code":BindManualCode,
		tablePrefix+"manual_code_operation_record":BindManualCodeOperationRecord,
		tablePrefix+"menu":BindMenu,
		tablePrefix+"menu_copy":BindMenuCopy,
		tablePrefix+"message":BindMessage,
		tablePrefix+"message_app":BindMessageApp,
		tablePrefix+"message_event":BindMessageEvent,
		tablePrefix+"message_event_user":BindMessageEventUser,
		tablePrefix+"message_extend":BindMessageExtend,
		tablePrefix+"message_img":BindMessageImg,
		tablePrefix+"message_interact":BindMessageInteract,
		tablePrefix+"message_interact_user":BindMessageInteractUser,
		tablePrefix+"message_jike":BindMessageJike,
		tablePrefix+"message_jike_user":BindMessageJikeUser,
		tablePrefix+"message_label":BindMessageLabel,
		tablePrefix+"message_lovecar":BindMessageLovecar,
		tablePrefix+"message_lovecar_user":BindMessageLovecarUser,
		tablePrefix+"message_record":BindMessageRecord,
		tablePrefix+"message_setting":BindMessageSetting,
		tablePrefix+"message_store":BindMessageStore,
		tablePrefix+"message_store_user":BindMessageStoreUser,
		tablePrefix+"message_type":BindMessageType,
		tablePrefix+"message_type_content":BindMessageTypeContent,
		tablePrefix+"message_type_copy":BindMessageTypeCopy,
		tablePrefix+"message_url":BindMessageUrl,
		tablePrefix+"msg_broker":BindMsgBroker,
		tablePrefix+"music":BindMusic,
		tablePrefix+"music1":BindMusic1,
		tablePrefix+"music_type":BindMusicType,
		tablePrefix+"node_punch_auth":BindNodePunchAuth,
		tablePrefix+"node_punch_confirm":BindNodePunchConfirm,
		tablePrefix+"onlineservice_mute":BindOnlineserviceMute,
		tablePrefix+"onlineservice_order":BindOnlineserviceOrder,
		tablePrefix+"operate_activity":BindOperateActivity,
		tablePrefix+"operate_activity_check":BindOperateActivityCheck,
		tablePrefix+"order":BindOrder,
		tablePrefix+"order_main":BindOrderMain,
		tablePrefix+"org_bank_card":BindOrgBankCard,
		tablePrefix+"org_money":BindOrgMoney,
		tablePrefix+"org_money_bill":BindOrgMoneyBill,
		tablePrefix+"org_money_recharge":BindOrgMoneyRecharge,
		tablePrefix+"org_money_withdraw":BindOrgMoneyWithdraw,
		tablePrefix+"organization":BindOrganization,
		tablePrefix+"organization_activity":BindOrganizationActivity,
		tablePrefix+"organization_car":BindOrganizationCar,
		tablePrefix+"organization_check":BindOrganizationCheck,
		tablePrefix+"organization_check_info":BindOrganizationCheckInfo,
		tablePrefix+"organization_conf":BindOrganizationConf,
		tablePrefix+"organization_conf_val":BindOrganizationConfVal,
		tablePrefix+"organization_copy":BindOrganizationCopy,
		tablePrefix+"organization_copy_copy":BindOrganizationCopyCopy,
		tablePrefix+"organization_focus":BindOrganizationFocus,
		tablePrefix+"organization_log":BindOrganizationLog,
		tablePrefix+"organization_order":BindOrganizationOrder,
		tablePrefix+"organization_orderpoints_record":BindOrganizationOrderpointsRecord,
		tablePrefix+"organization_privilege":BindOrganizationPrivilege,
		tablePrefix+"organization_privilege_shop":BindOrganizationPrivilegeShop,
		tablePrefix+"organization_privilege_store":BindOrganizationPrivilegeStore,
		tablePrefix+"organization_time":BindOrganizationTime,
		tablePrefix+"organization_type":BindOrganizationType,
		tablePrefix+"partner":BindPartner,
		tablePrefix+"partner_blacklist":BindPartnerBlacklist,
		tablePrefix+"partner_invite":BindPartnerInvite,
		tablePrefix+"partner_log":BindPartnerLog,
		tablePrefix+"partner_period":BindPartnerPeriod,
		tablePrefix+"partner_qrcode":BindPartnerQrcode,
		tablePrefix+"partner_register":BindPartnerRegister,
		tablePrefix+"partner_results":BindPartnerResults,
		tablePrefix+"partner_task":BindPartnerTask,
		tablePrefix+"pay_alipay_log":BindPayAlipayLog,
		tablePrefix+"pay_wx_log":BindPayWxLog,
		tablePrefix+"phone":BindPhone,
		tablePrefix+"photo":BindPhoto,
		tablePrefix+"points_data":BindPointsData,
		tablePrefix+"points_exceed":BindPointsExceed,
		tablePrefix+"points_exp_records":BindPointsExpRecords,
		tablePrefix+"points_exp_rules":BindPointsExpRules,
		tablePrefix+"points_goods":BindPointsGoods,
		tablePrefix+"points_goods_pay":BindPointsGoodsPay,
		tablePrefix+"points_log":BindPointsLog,
		tablePrefix+"points_product":BindPointsProduct,
		tablePrefix+"points_product_records":BindPointsProductRecords,
		tablePrefix+"points_ranking":BindPointsRanking,
		tablePrefix+"points_sale_rules":BindPointsSaleRules,
		tablePrefix+"praise":BindPraise,
		tablePrefix+"project_examine":BindProjectExamine,
		tablePrefix+"punch":BindPunch,
		tablePrefix+"purchase":BindPurchase,
		tablePrefix+"purchase_factory":BindPurchaseFactory,
		tablePrefix+"purchase_fail":BindPurchaseFail,
		tablePrefix+"purchase_img":BindPurchaseImg,
		tablePrefix+"purchase_records":BindPurchaseRecords,
		tablePrefix+"purchase_sms":BindPurchaseSms,
		tablePrefix+"purchase_user":BindPurchaseUser,
		tablePrefix+"push_sms_config":BindPushSmsConfig,
		tablePrefix+"questionnaire":BindQuestionnaire,
		tablePrefix+"questionnaire_answer":BindQuestionnaireAnswer,
		tablePrefix+"questionnaire_main":BindQuestionnaireMain,
		tablePrefix+"questionnaire_user":BindQuestionnaireUser,
		tablePrefix+"rabbit_msg":BindRabbitMsg,
		tablePrefix+"recommend":BindRecommend,
		tablePrefix+"recommend_auth":BindRecommendAuth,
		tablePrefix+"recommend_content":BindRecommendContent,
		tablePrefix+"recommend_detail_201905":BindRecommendDetail201905,
		tablePrefix+"recommend_detail_201906":BindRecommendDetail201906,
		tablePrefix+"recommend_detail_201907":BindRecommendDetail201907,
		tablePrefix+"recommend_detail_201908":BindRecommendDetail201908,
		tablePrefix+"recommend_detail_201909":BindRecommendDetail201909,
		tablePrefix+"recommend_detail_201910":BindRecommendDetail201910,
		tablePrefix+"recommend_detail_201911":BindRecommendDetail201911,
		tablePrefix+"recommend_detail_201912":BindRecommendDetail201912,
		tablePrefix+"recommend_org":BindRecommendOrg,
		tablePrefix+"recommend_position":BindRecommendPosition,
		tablePrefix+"recommend_position_config":BindRecommendPositionConfig,
		tablePrefix+"recommend_record":BindRecommendRecord,
		tablePrefix+"recommend_time":BindRecommendTime,
		tablePrefix+"recommend_time_extend":BindRecommendTimeExtend,
		tablePrefix+"recommend_tongji_day":BindRecommendTongjiDay,
		tablePrefix+"recommend_tongji_month":BindRecommendTongjiMonth,
		tablePrefix+"redpack":BindRedpack,
		tablePrefix+"redpack_area":BindRedpackArea,
		tablePrefix+"redpack_check_record":BindRedpackCheckRecord,
		tablePrefix+"redpack_content":BindRedpackContent,
		tablePrefix+"redpack_msg_record":BindRedpackMsgRecord,
		tablePrefix+"redpack_org":BindRedpackOrg,
		tablePrefix+"redpack_position":BindRedpackPosition,
		tablePrefix+"redpack_prize":BindRedpackPrize,
		tablePrefix+"redpack_record":BindRedpackRecord,
		tablePrefix+"redpack_record_temp":BindRedpackRecordTemp,
		tablePrefix+"redpack_time":BindRedpackTime,
		tablePrefix+"redpack_user_card":BindRedpackUserCard,
		tablePrefix+"redpack_user_card_record":BindRedpackUserCardRecord,
		tablePrefix+"region":BindRegion,
		tablePrefix+"report":BindReport,
		tablePrefix+"report_img":BindReportImg,
		tablePrefix+"report_record":BindReportRecord,
		tablePrefix+"ride":BindRide,
		tablePrefix+"ride_check":BindRideCheck,
		tablePrefix+"ride_check_notes":BindRideCheckNotes,
		tablePrefix+"ride_line":BindRideLine,
		tablePrefix+"ride_node":BindRideNode,
		tablePrefix+"road_book":BindRoadBook,
		tablePrefix+"road_book_line":BindRoadBookLine,
		tablePrefix+"road_book_line_content":BindRoadBookLineContent,
		tablePrefix+"road_book_meet":BindRoadBookMeet,
		tablePrefix+"road_book_recommend":BindRoadBookRecommend,
		tablePrefix+"roadbook_arouse":BindRoadbookArouse,
		tablePrefix+"roadbook_arouse_abnormal":BindRoadbookArouseAbnormal,
		tablePrefix+"roadbook_arouse_account":BindRoadbookArouseAccount,
		tablePrefix+"roadbook_arouse_attend":BindRoadbookArouseAttend,
		tablePrefix+"roadbook_arouse_attend_census":BindRoadbookArouseAttendCensus,
		tablePrefix+"roadbook_arouse_flow":BindRoadbookArouseFlow,
		tablePrefix+"roadbook_arouse_log":BindRoadbookArouseLog,
		tablePrefix+"roadbook_arouse_real_account":BindRoadbookArouseRealAccount,
		tablePrefix+"roadbook_arouse_related":BindRoadbookArouseRelated,
		tablePrefix+"roadbook_line":BindRoadbookLine,
		tablePrefix+"roadbook_merchant":BindRoadbookMerchant,
		tablePrefix+"roadbook_merchant_apply":BindRoadbookMerchantApply,
		tablePrefix+"roadbook_merchant_invite":BindRoadbookMerchantInvite,
		tablePrefix+"roadbook_node_punch_census":BindRoadbookNodePunchCensus,
		tablePrefix+"roulette_lottery":BindRouletteLottery,
		tablePrefix+"roulette_lottery_attend":BindRouletteLotteryAttend,
		tablePrefix+"roulette_lottery_award":BindRouletteLotteryAward,
		tablePrefix+"saleinfo_org":BindSaleinfoOrg,
		tablePrefix+"saler_star":BindSalerStar,
		tablePrefix+"search_log":BindSearchLog,
		tablePrefix+"series":BindSeries,
		tablePrefix+"service_city":BindServiceCity,
		tablePrefix+"sms_api_log":BindSmsApiLog,
		tablePrefix+"sms_template":BindSmsTemplate,
		tablePrefix+"sms_template_log":BindSmsTemplateLog,
		tablePrefix+"sms_total":BindSmsTotal,
		tablePrefix+"sms_verify_code":BindSmsVerifyCode,
		tablePrefix+"special":BindSpecial,
		tablePrefix+"spm_app_download_log":BindSpmAppDownloadLog,
		tablePrefix+"spm_baidureferer_log":BindSpmBaidurefererLog,
		tablePrefix+"spm_cate":BindSpmCate,
		tablePrefix+"spm_designation":BindSpmDesignation,
		tablePrefix+"spm_keywords":BindSpmKeywords,
		tablePrefix+"spm_log":BindSpmLog,
		tablePrefix+"spm_log_copy":BindSpmLogCopy,
		tablePrefix+"spm_parameter":BindSpmParameter,
		tablePrefix+"spm_place":BindSpmPlace,
		tablePrefix+"spm_promoter":BindSpmPromoter,
		tablePrefix+"staff":BindStaff,
		tablePrefix+"statistics":BindStatistics,
		tablePrefix+"supertopic_apply":BindSupertopicApply,
		tablePrefix+"tag_car":BindTagCar,
		tablePrefix+"tag_cart_specific":BindTagCartSpecific,
		tablePrefix+"tags":BindTags,
		tablePrefix+"tags_category":BindTagsCategory,
		tablePrefix+"tagscategory":BindTagscategory,
		tablePrefix+"thread_activity":BindThreadActivity,
		tablePrefix+"thread_activity_apply":BindThreadActivityApply,
		tablePrefix+"thread_activity_car":BindThreadActivityCar,
		tablePrefix+"thread_activity_master":BindThreadActivityMaster,
		tablePrefix+"thread_activity_scene":BindThreadActivityScene,
		tablePrefix+"thread_activity_type":BindThreadActivityType,
		tablePrefix+"thread_activity_user":BindThreadActivityUser,
		tablePrefix+"thread_ask":BindThreadAsk,
		tablePrefix+"thread_ask_answer":BindThreadAskAnswer,
		tablePrefix+"thread_ask_answer_follow":BindThreadAskAnswerFollow,
		tablePrefix+"thread_census_ranking":BindThreadCensusRanking,
		tablePrefix+"thread_cycle":BindThreadCycle,
		tablePrefix+"thread_delete_reason":BindThreadDeleteReason,
		tablePrefix+"thread_evaluate":BindThreadEvaluate,
		tablePrefix+"thread_evaluate_car":BindThreadEvaluateCar,
		tablePrefix+"thread_goods":BindThreadGoods,
		tablePrefix+"thread_goods_category":BindThreadGoodsCategory,
		tablePrefix+"thread_integral_log":BindThreadIntegralLog,
		tablePrefix+"thread_mastersay":BindThreadMastersay,
		tablePrefix+"thread_mastersay_car":BindThreadMastersayCar,
		tablePrefix+"thread_mood":BindThreadMood,
		tablePrefix+"thread_play":BindThreadPlay,
		tablePrefix+"thread_roadbook_annex":BindThreadRoadbookAnnex,
		tablePrefix+"thread_roadbook_line_been":BindThreadRoadbookLineBeen,
		tablePrefix+"thread_roadbook_line_related":BindThreadRoadbookLineRelated,
		tablePrefix+"thread_roadbook_note_user":BindThreadRoadbookNoteUser,
		tablePrefix+"thread_roadbook_punch_census":BindThreadRoadbookPunchCensus,
		tablePrefix+"thread_show":BindThreadShow,
		tablePrefix+"thread_topic":BindThreadTopic,
		tablePrefix+"thread_topic_access":BindThreadTopicAccess,
		tablePrefix+"thread_topic_apply":BindThreadTopicApply,
		tablePrefix+"thread_topic_fans":BindThreadTopicFans,
		tablePrefix+"thread_visible":BindThreadVisible,
		tablePrefix+"thread_word_column":BindThreadWordColumn,
		tablePrefix+"thread_words":BindThreadWords,
		tablePrefix+"ticket":BindTicket,
		tablePrefix+"ticket_api_log":BindTicketApiLog,
		tablePrefix+"ticket_apply_user":BindTicketApplyUser,
		tablePrefix+"ticket_check_record":BindTicketCheckRecord,
		tablePrefix+"ticket_checkout_auth":BindTicketCheckoutAuth,
		tablePrefix+"ticket_free":BindTicketFree,
		tablePrefix+"ticket_free_handout":BindTicketFreeHandout,
		tablePrefix+"ticket_nums_record":BindTicketNumsRecord,
		tablePrefix+"ticket_paper_record":BindTicketPaperRecord,
		tablePrefix+"ticket_qd":BindTicketQd,
		tablePrefix+"ticket_qd_sale":BindTicketQdSale,
		tablePrefix+"ticket_redpack":BindTicketRedpack,
		tablePrefix+"ticket_redpack_log":BindTicketRedpackLog,
		tablePrefix+"ticket_refund":BindTicketRefund,
		tablePrefix+"ticket_refund_log":BindTicketRefundLog,
		tablePrefix+"ticket_send_record":BindTicketSendRecord,
		tablePrefix+"ticket_third_party":BindTicketThirdParty,
		tablePrefix+"ticket_user":BindTicketUser,
		tablePrefix+"tmp_data":BindTmpData,
		tablePrefix+"topic_category":BindTopicCategory,
		tablePrefix+"topic_exp_log":BindTopicExpLog,
		tablePrefix+"topic_log":BindTopicLog,
		tablePrefix+"topic_scompere":BindTopicScompere,
		tablePrefix+"topic_signin":BindTopicSignin,
		tablePrefix+"traffic_statistics":BindTrafficStatistics,
		tablePrefix+"traffic_statistics_log":BindTrafficStatisticsLog,
		tablePrefix+"traffic_statistics_log_group":BindTrafficStatisticsLogGroup,
		tablePrefix+"twitter":BindTwitter,
		tablePrefix+"twitter_album":BindTwitterAlbum,
		tablePrefix+"twitter_source":BindTwitterSource,
		tablePrefix+"twitter_topic_access":BindTwitterTopicAccess,
		tablePrefix+"twitter_visibility":BindTwitterVisibility,
		tablePrefix+"upload_apk":BindUploadApk,
		tablePrefix+"user_bank_card":BindUserBankCard,
		tablePrefix+"user_behaviour":BindUserBehaviour,
		tablePrefix+"user_car":BindUserCar,
		tablePrefix+"user_follow":BindUserFollow,
		tablePrefix+"user_points":BindUserPoints,
		tablePrefix+"user_recommend":BindUserRecommend,
		tablePrefix+"user_setting":BindUserSetting,
		tablePrefix+"video":BindVideo,
		tablePrefix+"video_attach":BindVideoAttach,
		tablePrefix+"video_attach_receive":BindVideoAttachReceive,
		tablePrefix+"video_log":BindVideoLog,
		tablePrefix+"video_notice":BindVideoNotice,
		tablePrefix+"video_recommend_table":BindVideoRecommendTable,
		tablePrefix+"video_visible":BindVideoVisible,
		tablePrefix+"virtual_nickname":BindVirtualNickname,
		tablePrefix+"virtual_record":BindVirtualRecord,
		tablePrefix+"wanna_buy":BindWannaBuy,
		tablePrefix+"weather_citys":BindWeatherCitys,
		tablePrefix+"weather_today":BindWeatherToday,
		tablePrefix+"weixin_config":BindWeixinConfig,
		tablePrefix+"weixin_custom_send":BindWeixinCustomSend,
		tablePrefix+"weixin_kerword":BindWeixinKerword,
		tablePrefix+"weixin_material":BindWeixinMaterial,
		tablePrefix+"weixin_menu":BindWeixinMenu,
		tablePrefix+"weixin_push":BindWeixinPush,
		tablePrefix+"weixin_push_auto":BindWeixinPushAuto,
		tablePrefix+"weixin_push_config":BindWeixinPushConfig,
		tablePrefix+"weixin_push_interactive":BindWeixinPushInteractive,
		tablePrefix+"weixin_push_material":BindWeixinPushMaterial,
		tablePrefix+"weixin_push_record":BindWeixinPushRecord,
		tablePrefix+"weixin_push_tag":BindWeixinPushTag,
		tablePrefix+"weixin_qrcode":BindWeixinQrcode,
		tablePrefix+"weixin_qrcode_scan_log":BindWeixinQrcodeScanLog,
		tablePrefix+"weixin_template":BindWeixinTemplate,
		tablePrefix+"xc1_jurisdiction":BindXc1Jurisdiction,
		tablePrefix+"xc1_manage_object":BindXc1ManageObject,
		tablePrefix+"xc1_org2scene":BindXc1Org2scene,
		tablePrefix+"xc1_worker2guest":BindXc1Worker2guest,
		tablePrefix+"xc_app_message":BindXcAppMessage,
		tablePrefix+"xc_app_message_count":BindXcAppMessageCount,
		tablePrefix+"xc_apply_item":BindXcApplyItem,
		tablePrefix+"xc_apply_type":BindXcApplyType,
		tablePrefix+"xc_area":BindXcArea,
		tablePrefix+"xc_auth_mapping":BindXcAuthMapping,
		tablePrefix+"xc_authorise":BindXcAuthorise,
		tablePrefix+"xc_check_course":BindXcCheckCourse,
		tablePrefix+"xc_check_table":BindXcCheckTable,
		tablePrefix+"xc_checker2object":BindXcChecker2object,
		tablePrefix+"xc_friends":BindXcFriends,
		tablePrefix+"xc_guest2object":BindXcGuest2object,
		tablePrefix+"xc_job_type":BindXcJobType,
		tablePrefix+"xc_message":BindXcMessage,
		tablePrefix+"xc_message_addtask":BindXcMessageAddtask,
		tablePrefix+"xc_message_apply":BindXcMessageApply,
		tablePrefix+"xc_message_content":BindXcMessageContent,
		tablePrefix+"xc_message_rectification":BindXcMessageRectification,
		tablePrefix+"xc_message_setting":BindXcMessageSetting,
		tablePrefix+"xc_message_unqualifiedcheck":BindXcMessageUnqualifiedcheck,
		tablePrefix+"xc_object":BindXcObject,
		tablePrefix+"xc_object2dispose":BindXcObject2dispose,
		tablePrefix+"xc_object_organization":BindXcObjectOrganization,
		tablePrefix+"xc_object_type":BindXcObjectType,
		tablePrefix+"xc_offset_trigger_time":BindXcOffsetTriggerTime,
		tablePrefix+"xc_order":BindXcOrder,
		tablePrefix+"xc_order_progress":BindXcOrderProgress,
		tablePrefix+"xc_organization":BindXcOrganization,
		tablePrefix+"xc_pay_log":BindXcPayLog,
		tablePrefix+"xc_permission":BindXcPermission,
		tablePrefix+"xc_scene":BindXcScene,
		tablePrefix+"xc_scene_review":BindXcSceneReview,
		tablePrefix+"xc_staff2object":BindXcStaff2object,
		tablePrefix+"xc_task":BindXcTask,
		tablePrefix+"xc_task_copy":BindXcTaskCopy,
		tablePrefix+"xc_task_dis":BindXcTaskDis,
		tablePrefix+"xc_task_disimg":BindXcTaskDisimg,
		tablePrefix+"xc_task_log":BindXcTaskLog,
		tablePrefix+"xc_task_package":BindXcTaskPackage,
		tablePrefix+"xc_top_up_log":BindXcTopUpLog,
		tablePrefix+"xc_user2scene":BindXcUser2scene,
		tablePrefix+"xc_usercoordinate":BindXcUsercoordinate,
		tablePrefix+"xc_worker":BindXcWorker,
		tablePrefix+"xc_worker2area":BindXcWorker2area,
		tablePrefix+"xc_worker_yanghuan_back":BindXcWorkerYanghuanBack,
		tablePrefix+"yaoyiyao":BindYaoyiyao,
		tablePrefix+"yaoyiyao_auth":BindYaoyiyaoAuth,
		tablePrefix+"yaoyiyao_awards":BindYaoyiyaoAwards,
		tablePrefix+"yaoyiyao_awards_copy":BindYaoyiyaoAwardsCopy,
		tablePrefix+"yaoyiyao_cacheawards":BindYaoyiyaoCacheawards,
		tablePrefix+"yaoyiyao_end_log":BindYaoyiyaoEndLog,
		tablePrefix+"yaoyiyao_not_win":BindYaoyiyaoNotWin,
		tablePrefix+"yaoyiyao_pass_log":BindYaoyiyaoPassLog,
		tablePrefix+"yaoyiyao_rank":BindYaoyiyaoRank,
		tablePrefix+"yaoyiyao_rotary":BindYaoyiyaoRotary,
		tablePrefix+"yaoyiyao_signin":BindYaoyiyaoSignin,
		tablePrefix+"yaoyiyao_status":BindYaoyiyaoStatus,
	}


//绑定thread
func BindThread(args Args) interface{} {

	model := models.TgThread{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

//绑定thread_roadbook
func Bind2handcarApply(args Args) interface{} {

	model := models.Tg2handcarApply{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}
//绑定thread_roadbook
func BindThreadRoadbook(args Args) interface{} {

	model := models.TgThreadRoadbook{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func Bind2handcarBooth(args Args) interface{} {

	model := models.Tg2handcarBooth{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func Bind2handcarComplain(args Args) interface{} {

	model := models.Tg2handcarComplain{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func Bind2handcarConfig(args Args) interface{} {

	model := models.Tg2handcarConfig{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func Bind2handcarDebitList(args Args) interface{} {

	model := models.Tg2handcarDebitList{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func Bind2handcarIntention(args Args) interface{} {

	model := models.Tg2handcarIntention{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func Bind2handcarList(args Args) interface{} {

	model := models.Tg2handcarList{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func Bind2handcarOrder(args Args) interface{} {

	model := models.Tg2handcarOrder{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func Bind2handcarPark(args Args) interface{} {

	model := models.Tg2handcarPark{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func Bind2handcarParkTime(args Args) interface{} {

	model := models.Tg2handcarParkTime{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func Bind2handcarParkType(args Args) interface{} {

	model := models.Tg2handcarParkType{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func Bind2handcarPlan(args Args) interface{} {

	model := models.Tg2handcarPlan{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func Bind2handcarTime(args Args) interface{} {

	model := models.Tg2handcarTime{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func Bind2handcarWxFail(args Args) interface{} {

	model := models.Tg2handcarWxFail{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindAccountManage(args Args) interface{} {

	model := models.TgAccountManage{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindActionLog(args Args) interface{} {

	model := models.TgActionLog{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindActivityArea(args Args) interface{} {

	model := models.TgActivityArea{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindActivityForward(args Args) interface{} {

	model := models.TgActivityForward{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindActivityForwardContent(args Args) interface{} {

	model := models.TgActivityForwardContent{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindActivityForwardPrize(args Args) interface{} {

	model := models.TgActivityForwardPrize{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindActivityForwardRecord(args Args) interface{} {

	model := models.TgActivityForwardRecord{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindActivityForwardRecordDetails(args Args) interface{} {

	model := models.TgActivityForwardRecordDetails{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindActivityLine(args Args) interface{} {

	model := models.TgActivityLine{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindActivityLineNodeRelated(args Args) interface{} {

	model := models.TgActivityLineNodeRelated{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindActivityLotteryAttendance(args Args) interface{} {

	model := models.TgActivityLotteryAttendance{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindActivityLotteryAuth(args Args) interface{} {

	model := models.TgActivityLotteryAuth{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindActivityLotteryCode(args Args) interface{} {

	model := models.TgActivityLotteryCode{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindActivityLotteryPrize(args Args) interface{} {

	model := models.TgActivityLotteryPrize{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindActivityLotteryRotary(args Args) interface{} {

	model := models.TgActivityLotteryRotary{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindActivityLotterySms(args Args) interface{} {

	model := models.TgActivityLotterySms{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindActivityLotteryStatus(args Args) interface{} {

	model := models.TgActivityLotteryStatus{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindActivityTable(args Args) interface{} {

	model := models.TgActivityTable{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindActivityTags(args Args) interface{} {

	model := models.TgActivityTags{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindActivityTurntable(args Args) interface{} {

	model := models.TgActivityTurntable{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindActivityTurntableContent(args Args) interface{} {

	model := models.TgActivityTurntableContent{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindActivityTurntablePrize(args Args) interface{} {

	model := models.TgActivityTurntablePrize{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindActivityTurntableQrcodetag(args Args) interface{} {

	model := models.TgActivityTurntableQrcodetag{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindActivityTurntableRecord(args Args) interface{} {

	model := models.TgActivityTurntableRecord{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindActivityTurntableResult(args Args) interface{} {

	model := models.TgActivityTurntableResult{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindActivityTurntableSaddress(args Args) interface{} {

	model := models.TgActivityTurntableSaddress{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindActivityTurntableShare(args Args) interface{} {

	model := models.TgActivityTurntableShare{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindActivityTurntableUser(args Args) interface{} {

	model := models.TgActivityTurntableUser{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindActivityVoteReferOpt(args Args) interface{} {

	model := models.TgActivityVoteReferOpt{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindActivityVoteRules(args Args) interface{} {

	model := models.TgActivityVoteRules{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindActivityVoteUserLimit(args Args) interface{} {

	model := models.TgActivityVoteUserLimit{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindActivityVoteUserTable(args Args) interface{} {

	model := models.TgActivityVoteUserTable{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindActivityVoteUserTicketcert(args Args) interface{} {

	model := models.TgActivityVoteUserTicketcert{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindAdminAuthProjectXc(args Args) interface{} {

	model := models.TgAdminAuthProjectXc{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindAdminAuthProjectXcCopy(args Args) interface{} {

	model := models.TgAdminAuthProjectXcCopy{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindAdminAuthUser(args Args) interface{} {

	model := models.TgAdminAuthUser{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindAdminAuthUserCopy(args Args) interface{} {

	model := models.TgAdminAuthUserCopy{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindAdminMenu(args Args) interface{} {

	model := models.TgAdminMenu{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindAdminReport(args Args) interface{} {

	model := models.TgAdminReport{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindAdminRole(args Args) interface{} {

	model := models.TgAdminRole{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindAdminUser(args Args) interface{} {

	model := models.TgAdminUser{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindAdminUser2area(args Args) interface{} {

	model := models.TgAdminUser2area{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindAdminUser2areaOrg(args Args) interface{} {

	model := models.TgAdminUser2areaOrg{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindAdminUser2org(args Args) interface{} {

	model := models.TgAdminUser2org{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindAdminUser2orgCopy(args Args) interface{} {

	model := models.TgAdminUser2orgCopy{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindAdminUser2orgSetting(args Args) interface{} {

	model := models.TgAdminUser2orgSetting{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindAdminUserAlltags(args Args) interface{} {

	model := models.TgAdminUserAlltags{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindAdminUserAuth(args Args) interface{} {

	model := models.TgAdminUserAuth{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindAdminUserAuthCopy(args Args) interface{} {

	model := models.TgAdminUserAuthCopy{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindAdminUserAvatar(args Args) interface{} {

	model := models.TgAdminUserAvatar{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindAdminUserCopy1(args Args) interface{} {

	model := models.TgAdminUserCopy1{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindAdminUserDepartment(args Args) interface{} {

	model := models.TgAdminUserDepartment{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindAdminUserDoyen(args Args) interface{} {

	model := models.TgAdminUserDoyen{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindAdminUserFavor(args Args) interface{} {

	model := models.TgAdminUserFavor{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindAdminUserImpression(args Args) interface{} {

	model := models.TgAdminUserImpression{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindAdminUserLock(args Args) interface{} {

	model := models.TgAdminUserLock{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindAdminUserLoginFail(args Args) interface{} {

	model := models.TgAdminUserLoginFail{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindAdminUserMoney(args Args) interface{} {

	model := models.TgAdminUserMoney{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindAdminUserMoneyBill(args Args) interface{} {

	model := models.TgAdminUserMoneyBill{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindAdminUserMoneyRecharge(args Args) interface{} {

	model := models.TgAdminUserMoneyRecharge{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindAdminUserMoneyWithdraw(args Args) interface{} {

	model := models.TgAdminUserMoneyWithdraw{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindAdminUserOrg(args Args) interface{} {

	model := models.TgAdminUserOrg{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindAdminUserPosition(args Args) interface{} {

	model := models.TgAdminUserPosition{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindAdminUserPunish(args Args) interface{} {

	model := models.TgAdminUserPunish{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindAdminUserRealNameAuth(args Args) interface{} {

	model := models.TgAdminUserRealNameAuth{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindAdminUserRole(args Args) interface{} {

	model := models.TgAdminUserRole{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindAdminUserTags(args Args) interface{} {

	model := models.TgAdminUserTags{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindAdminUserType(args Args) interface{} {

	model := models.TgAdminUserType{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindAdvertisementAdminLog(args Args) interface{} {

	model := models.TgAdvertisementAdminLog{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindAdvertisementContent(args Args) interface{} {

	model := models.TgAdvertisementContent{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindAdvertisementOrgStrategy(args Args) interface{} {

	model := models.TgAdvertisementOrgStrategy{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindAdvertisementPosition(args Args) interface{} {

	model := models.TgAdvertisementPosition{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindAgreement(args Args) interface{} {

	model := models.TgAgreement{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindAgreementPlace(args Args) interface{} {

	model := models.TgAgreementPlace{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindAlbum(args Args) interface{} {

	model := models.TgAlbum{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindAlltags(args Args) interface{} {

	model := models.TgAlltags{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindAlltagsArticle(args Args) interface{} {

	model := models.TgAlltagsArticle{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindAlltagsCategory(args Args) interface{} {

	model := models.TgAlltagsCategory{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindAlltagsCategoryActivity(args Args) interface{} {

	model := models.TgAlltagsCategoryActivity{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindArea(args Args) interface{} {

	model := models.TgArea{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindAreaBak(args Args) interface{} {

	model := models.TgAreaBak{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindAreaHot(args Args) interface{} {

	model := models.TgAreaHot{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}


func BindArticle(args Args) interface{} {

	model := models.TgArticle{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindArticleCategory(args Args) interface{} {

	model := models.TgArticleCategory{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindArticleData(args Args) interface{} {

	model := models.TgArticleData{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindArticleTag(args Args) interface{} {

	model := models.TgArticleTag{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindAskCate(args Args) interface{} {

	model := models.TgAskCate{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindAskExpert(args Args) interface{} {

	model := models.TgAskExpert{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindAsyncTask(args Args) interface{} {

	model := models.TgAsyncTask{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindAttachment(args Args) interface{} {

	model := models.TgAttachment{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindAuthAction(args Args) interface{} {

	model := models.TgAuthAction{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindAuthRole(args Args) interface{} {

	model := models.TgAuthRole{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindAuthRoleAction(args Args) interface{} {

	model := models.TgAuthRoleAction{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindAuthRoleSource(args Args) interface{} {

	model := models.TgAuthRoleSource{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindAuthUserRole(args Args) interface{} {

	model := models.TgAuthUserRole{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindBaiduConfig(args Args) interface{} {

	model := models.TgBaiduConfig{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindBaiduNotice(args Args) interface{} {

	model := models.TgBaiduNotice{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindBaiduShare(args Args) interface{} {

	model := models.TgBaiduShare{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindBenefit(args Args) interface{} {

	model := models.TgBenefit{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindBenefitAccess(args Args) interface{} {

	model := models.TgBenefitAccess{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindBenefitStore(args Args) interface{} {

	model := models.TgBenefitStore{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindBidKeyword(args Args) interface{} {

	model := models.TgBidKeyword{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindBidKeywordCopy(args Args) interface{} {

	model := models.TgBidKeywordCopy{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindBrand(args Args) interface{} {

	model := models.TgBrand{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindBrowseHistory(args Args) interface{} {

	model := models.TgBrowseHistory{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindCallAudience(args Args) interface{} {

	model := models.TgCallAudience{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindCallAudienceAccess(args Args) interface{} {

	model := models.TgCallAudienceAccess{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindCallAudienceConf(args Args) interface{} {

	model := models.TgCallAudienceConf{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindCallAudienceConfVal(args Args) interface{} {

	model := models.TgCallAudienceConfVal{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindCallAudienceReview(args Args) interface{} {

	model := models.TgCallAudienceReview{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindCallAudienceReviewLog(args Args) interface{} {

	model := models.TgCallAudienceReviewLog{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindCallAudienceSmstask(args Args) interface{} {

	model := models.TgCallAudienceSmstask{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindCallAudienceSmstaskContent(args Args) interface{} {

	model := models.TgCallAudienceSmstaskContent{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindCallAudienceSmstaskRecord(args Args) interface{} {

	model := models.TgCallAudienceSmstaskRecord{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindCallAudienceUser(args Args) interface{} {

	model := models.TgCallAudienceUser{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindCallAudienceVisit(args Args) interface{} {

	model := models.TgCallAudienceVisit{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindCallDatatype(args Args) interface{} {

	model := models.TgCallDatatype{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindCallDraw(args Args) interface{} {

	model := models.TgCallDraw{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindCallEventAuth(args Args) interface{} {

	model := models.TgCallEventAuth{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindCallPrivacy(args Args) interface{} {

	model := models.TgCallPrivacy{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindCallQuestionnaire(args Args) interface{} {

	model := models.TgCallQuestionnaire{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindCallQuestionnaireConf(args Args) interface{} {

	model := models.TgCallQuestionnaireConf{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindCallQuestionnaireUser(args Args) interface{} {

	model := models.TgCallQuestionnaireUser{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindCallQuestionnaireVal(args Args) interface{} {

	model := models.TgCallQuestionnaireVal{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindCallRecord(args Args) interface{} {

	model := models.TgCallRecord{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindCallUserAuth(args Args) interface{} {

	model := models.TgCallUserAuth{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindCar(args Args) interface{} {

	model := models.TgCar{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindCar2conf(args Args) interface{} {

	model := models.TgCar2conf{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindCarColor(args Args) interface{} {

	model := models.TgCarColor{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindCarConf(args Args) interface{} {

	model := models.TgCarConf{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindCarConfVal(args Args) interface{} {

	model := models.TgCarConfVal{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindCarCopy(args Args) interface{} {

	model := models.TgCarCopy{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindCarCustomMade(args Args) interface{} {

	model := models.TgCarCustomMade{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindCarFind(args Args) interface{} {

	model := models.TgCarFind{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindCarLove(args Args) interface{} {

	model := models.TgCarLove{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindCarOrder(args Args) interface{} {

	model := models.TgCarOrder{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindCarOrderConfig(args Args) interface{} {

	model := models.TgCarOrderConfig{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindCarOrderList(args Args) interface{} {

	model := models.TgCarOrderList{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindCarOrderMsg(args Args) interface{} {

	model := models.TgCarOrderMsg{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindCarOrderOffer(args Args) interface{} {

	model := models.TgCarOrderOffer{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindCarOrderOfferCount(args Args) interface{} {

	model := models.TgCarOrderOfferCount{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindCarOrderPushList(args Args) interface{} {

	model := models.TgCarOrderPushList{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindCarOrderTags(args Args) interface{} {

	model := models.TgCarOrderTags{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindCarPic(args Args) interface{} {

	model := models.TgCarPic{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindCarSearchLog(args Args) interface{} {

	model := models.TgCarSearchLog{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindCarSeries(args Args) interface{} {

	model := models.TgCarSeries{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindCarSeriesColor(args Args) interface{} {

	model := models.TgCarSeriesColor{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindCarorder(args Args) interface{} {

	model := models.TgCarorder{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindCarorderConfig(args Args) interface{} {

	model := models.TgCarorderConfig{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindCarorderList(args Args) interface{} {

	model := models.TgCarorderList{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindCarorderLovecar(args Args) interface{} {

	model := models.TgCarorderLovecar{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindCarorderOffer(args Args) interface{} {

	model := models.TgCarorderOffer{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindCarorderOfferRecord(args Args) interface{} {

	model := models.TgCarorderOfferRecord{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindCarorderPushList(args Args) interface{} {

	model := models.TgCarorderPushList{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindCarorderRecord(args Args) interface{} {

	model := models.TgCarorderRecord{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindCarorderTactics(args Args) interface{} {

	model := models.TgCarorderTactics{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindCarorderTacticsConfig(args Args) interface{} {

	model := models.TgCarorderTacticsConfig{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindCarorderTacticsGift(args Args) interface{} {

	model := models.TgCarorderTacticsGift{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindCarorderTacticsRecord(args Args) interface{} {

	model := models.TgCarorderTacticsRecord{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindCarousel(args Args) interface{} {

	model := models.TgCarousel{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindCarouselItem(args Args) interface{} {

	model := models.TgCarouselItem{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindCarouselObject(args Args) interface{} {

	model := models.TgCarouselObject{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindCarouselOrgPlan(args Args) interface{} {

	model := models.TgCarouselOrgPlan{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindCarouselOrgRight(args Args) interface{} {

	model := models.TgCarouselOrgRight{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindCarouselOriginality(args Args) interface{} {

	model := models.TgCarouselOriginality{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindCarouselScheduling(args Args) interface{} {

	model := models.TgCarouselScheduling{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindCarouselSchedulingAuthorize(args Args) interface{} {

	model := models.TgCarouselSchedulingAuthorize{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindCarouselStrategy(args Args) interface{} {

	model := models.TgCarouselStrategy{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindCarouselSubjectItem(args Args) interface{} {

	model := models.TgCarouselSubjectItem{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindCert(args Args) interface{} {

	model := models.TgCert{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindCertApply(args Args) interface{} {

	model := models.TgCertApply{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindCertContent(args Args) interface{} {

	model := models.TgCertContent{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindCertHandout(args Args) interface{} {

	model := models.TgCertHandout{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindCertNumsRecord(args Args) interface{} {

	model := models.TgCertNumsRecord{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindCertUser(args Args) interface{} {

	model := models.TgCertUser{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindCertUserCopy(args Args) interface{} {

	model := models.TgCertUserCopy{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindCertUserRecord(args Args) interface{} {

	model := models.TgCertUserRecord{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindCertUserRecordAct(args Args) interface{} {

	model := models.TgCertUserRecordAct{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindChickenSoup(args Args) interface{} {

	model := models.TgChickenSoup{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindCollect(args Args) interface{} {

	model := models.TgCollect{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindCollectMapping(args Args) interface{} {

	model := models.TgCollectMapping{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindComment(args Args) interface{} {

	model := models.TgComment{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindCommentLog(args Args) interface{} {

	model := models.TgCommentLog{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindConfig(args Args) interface{} {

	model := models.TgConfig{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindCountry(args Args) interface{} {

	model := models.TgCountry{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindCoupon(args Args) interface{} {

	model := models.TgCoupon{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindCouponCar(args Args) interface{} {

	model := models.TgCouponCar{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindCouponCategory(args Args) interface{} {

	model := models.TgCouponCategory{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindCouponCheck(args Args) interface{} {

	model := models.TgCouponCheck{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindCouponCode(args Args) interface{} {

	model := models.TgCouponCode{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindCouponDealers(args Args) interface{} {

	model := models.TgCouponDealers{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindCouponDistribute(args Args) interface{} {

	model := models.TgCouponDistribute{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindCouponDistributeList(args Args) interface{} {

	model := models.TgCouponDistributeList{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindCouponDistributeListReissue(args Args) interface{} {

	model := models.TgCouponDistributeListReissue{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindCouponDistributeReissue(args Args) interface{} {

	model := models.TgCouponDistributeReissue{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindCouponIssue(args Args) interface{} {

	model := models.TgCouponIssue{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindCouponOrder(args Args) interface{} {

	model := models.TgCouponOrder{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindCouponPush(args Args) interface{} {

	model := models.TgCouponPush{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindCouponPushLog(args Args) interface{} {

	model := models.TgCouponPushLog{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindCouponRefund(args Args) interface{} {

	model := models.TgCouponRefund{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindCouponReplace(args Args) interface{} {

	model := models.TgCouponReplace{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindCouponReplaceDistribute(args Args) interface{} {

	model := models.TgCouponReplaceDistribute{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindCouponReplaceReissue(args Args) interface{} {

	model := models.TgCouponReplaceReissue{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindCouponShareLog(args Args) interface{} {

	model := models.TgCouponShareLog{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindCouponTicket(args Args) interface{} {

	model := models.TgCouponTicket{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindCouponWarehouse(args Args) interface{} {

	model := models.TgCouponWarehouse{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindCouponWarehouseCode(args Args) interface{} {

	model := models.TgCouponWarehouseCode{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindCrawlerData(args Args) interface{} {

	model := models.TgCrawlerData{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindDanger(args Args) interface{} {

	model := models.TgDanger{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindDshowBanner(args Args) interface{} {

	model := models.TgDshowBanner{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindDshowBase(args Args) interface{} {

	model := models.TgDshowBase{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindDshowGift(args Args) interface{} {

	model := models.TgDshowGift{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindDshowShare(args Args) interface{} {

	model := models.TgDshowShare{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindDshowText(args Args) interface{} {

	model := models.TgDshowText{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindEasemobMessages(args Args) interface{} {

	model := models.TgEasemobMessages{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindEasemobMessagesNum(args Args) interface{} {

	model := models.TgEasemobMessagesNum{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindEasemobMessagesPayload(args Args) interface{} {

	model := models.TgEasemobMessagesPayload{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindEmoji(args Args) interface{} {

	model := models.TgEmoji{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindEmojiGroup(args Args) interface{} {

	model := models.TgEmojiGroup{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindEmojiUserAccess(args Args) interface{} {

	model := models.TgEmojiUserAccess{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindEntranceWay(args Args) interface{} {

	model := models.TgEntranceWay{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindEvent(args Args) interface{} {

	model := models.TgEvent{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindEventActivity(args Args) interface{} {

	model := models.TgEventActivity{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindEventActivityData(args Args) interface{} {

	model := models.TgEventActivityData{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindEventAd(args Args) interface{} {

	model := models.TgEventAd{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindEventBdconfig(args Args) interface{} {

	model := models.TgEventBdconfig{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindEventBooth(args Args) interface{} {

	model := models.TgEventBooth{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindEventBrand(args Args) interface{} {

	model := models.TgEventBrand{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindEventBrandMerge(args Args) interface{} {

	model := models.TgEventBrandMerge{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindEventBusiness(args Args) interface{} {

	model := models.TgEventBusiness{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindEventDescImg(args Args) interface{} {

	model := models.TgEventDescImg{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindEventExamine(args Args) interface{} {

	model := models.TgEventExamine{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindEventGuest(args Args) interface{} {

	model := models.TgEventGuest{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindEventHall(args Args) interface{} {

	model := models.TgEventHall{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindEventImgConfig(args Args) interface{} {

	model := models.TgEventImgConfig{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindEventImgUnshow(args Args) interface{} {

	model := models.TgEventImgUnshow{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindEventInformation(args Args) interface{} {

	model := models.TgEventInformation{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindEventModel(args Args) interface{} {

	model := models.TgEventModel{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindEventNavigation(args Args) interface{} {

	model := models.TgEventNavigation{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindEventNavigationConf(args Args) interface{} {

	model := models.TgEventNavigationConf{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindEventNewcar(args Args) interface{} {

	model := models.TgEventNewcar{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindEventSchedule(args Args) interface{} {

	model := models.TgEventSchedule{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindEventShowcar(args Args) interface{} {

	model := models.TgEventShowcar{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindEventStaff(args Args) interface{} {

	model := models.TgEventStaff{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindEventStaffAuth(args Args) interface{} {

	model := models.TgEventStaffAuth{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindEventStaffAuthArea(args Args) interface{} {

	model := models.TgEventStaffAuthArea{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindEventStaffAuthBrand(args Args) interface{} {

	model := models.TgEventStaffAuthBrand{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindEventStaffAuthUser(args Args) interface{} {

	model := models.TgEventStaffAuthUser{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindEventUser(args Args) interface{} {

	model := models.TgEventUser{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindFab(args Args) interface{} {

	model := models.TgFab{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindFactory(args Args) interface{} {

	model := models.TgFactory{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindFailedJobs(args Args) interface{} {

	model := models.TgFailedJobs{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindFeedback(args Args) interface{} {

	model := models.TgFeedback{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindFeedsIgnore(args Args) interface{} {

	model := models.TgFeedsIgnore{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindFeedsTopic(args Args) interface{} {

	model := models.TgFeedsTopic{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindFeedsUser(args Args) interface{} {

	model := models.TgFeedsUser{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindFollowGroup(args Args) interface{} {

	model := models.TgFollowGroup{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindFriend(args Args) interface{} {

	model := models.TgFriend{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindFriendBuyinfo(args Args) interface{} {

	model := models.TgFriendBuyinfo{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindFriendGroup(args Args) interface{} {

	model := models.TgFriendGroup{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindFriendGroupAccess(args Args) interface{} {

	model := models.TgFriendGroupAccess{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindFriendLatent(args Args) interface{} {

	model := models.TgFriendLatent{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindFriendManageLog(args Args) interface{} {

	model := models.TgFriendManageLog{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindFriendOrg(args Args) interface{} {

	model := models.TgFriendOrg{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindGallery(args Args) interface{} {

	model := models.TgGallery{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindGroup(args Args) interface{} {

	model := models.TgGroup{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindGroupAuthentication(args Args) interface{} {

	model := models.TgGroupAuthentication{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindGroupBlacklist(args Args) interface{} {

	model := models.TgGroupBlacklist{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindGroupCategory(args Args) interface{} {

	model := models.TgGroupCategory{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindGroupInvitation(args Args) interface{} {

	model := models.TgGroupInvitation{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindGroupNotification(args Args) interface{} {

	model := models.TgGroupNotification{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindGroupNotificationRead(args Args) interface{} {

	model := models.TgGroupNotificationRead{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindGroupQrcode(args Args) interface{} {

	model := models.TgGroupQrcode{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindGroupTag(args Args) interface{} {

	model := models.TgGroupTag{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindGroupTagAccess(args Args) interface{} {

	model := models.TgGroupTagAccess{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindGroupUser(args Args) interface{} {

	model := models.TgGroupUser{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindGroupUserKickOut(args Args) interface{} {

	model := models.TgGroupUserKickOut{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindHelpQuestion(args Args) interface{} {

	model := models.TgHelpQuestion{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindHotSearch(args Args) interface{} {

	model := models.TgHotSearch{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}


func BindHotSearchExt(args Args) interface{} {

	model := models.TgHotSearchExt{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindImg(args Args) interface{} {

	model := models.TgImg{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindImgTag(args Args) interface{} {

	model := models.TgImgTag{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindImgType(args Args) interface{} {

	model := models.TgImgType{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindIndustry(args Args) interface{} {

	model := models.TgIndustry{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindIndustryUser(args Args) interface{} {

	model := models.TgIndustryUser{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}


func BindKoubei(args Args) interface{} {

	model := models.TgKoubei{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindKoubeiAlbum(args Args) interface{} {

	model := models.TgKoubeiAlbum{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindKoubeiSpider(args Args) interface{} {

	model := models.TgKoubeiSpider{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindLatentAutoset(args Args) interface{} {

	model := models.TgLatentAutoset{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindLatentDispense(args Args) interface{} {

	model := models.TgLatentDispense{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindLine(args Args) interface{} {

	model := models.TgLine{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindLineNode(args Args) interface{} {

	model := models.TgLineNode{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindLineNodeAlbum(args Args) interface{} {

	model := models.TgLineNodeAlbum{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindLineNodeComment(args Args) interface{} {

	model := models.TgLineNodeComment{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindLineNodeCommentMinutia(args Args) interface{} {

	model := models.TgLineNodeCommentMinutia{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindLineNodeCommentText(args Args) interface{} {

	model := models.TgLineNodeCommentText{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindLineNodePunchCensus(args Args) interface{} {

	model := models.TgLineNodePunchCensus{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindLineNodeRelated(args Args) interface{} {

	model := models.TgLineNodeRelated{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindLiveMessage(args Args) interface{} {

	model := models.TgLiveMessage{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindLiveUser(args Args) interface{} {

	model := models.TgLiveUser{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindLogAdminLogin(args Args) interface{} {

	model := models.TgLogAdminLogin{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindLogAdminWork(args Args) interface{} {

	model := models.TgLogAdminWork{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindLottery(args Args) interface{} {

	model := models.TgLottery{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindLotteryAuth(args Args) interface{} {

	model := models.TgLotteryAuth{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindLotteryConfig(args Args) interface{} {

	model := models.TgLotteryConfig{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindLotteryCron(args Args) interface{} {

	model := models.TgLotteryCron{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindLotteryNew(args Args) interface{} {

	model := models.TgLotteryNew{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindLotteryNewAuth(args Args) interface{} {

	model := models.TgLotteryNewAuth{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindLotteryNewCodeBill(args Args) interface{} {

	model := models.TgLotteryNewCodeBill{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindLotteryNewConfig(args Args) interface{} {

	model := models.TgLotteryNewConfig{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindLotteryNewCron(args Args) interface{} {

	model := models.TgLotteryNewCron{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindLotteryNewLuckyCode(args Args) interface{} {

	model := models.TgLotteryNewLuckyCode{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindLotteryNewOrder(args Args) interface{} {

	model := models.TgLotteryNewOrder{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindLotteryNewOrgCode(args Args) interface{} {

	model := models.TgLotteryNewOrgCode{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindLotteryNewPrize(args Args) interface{} {

	model := models.TgLotteryNewPrize{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindLotteryNewRecord(args Args) interface{} {

	model := models.TgLotteryNewRecord{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindLotteryNewSet(args Args) interface{} {

	model := models.TgLotteryNewSet{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindLotteryNewUser(args Args) interface{} {

	model := models.TgLotteryNewUser{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindLotteryNewUserCode(args Args) interface{} {

	model := models.TgLotteryNewUserCode{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindLotteryOrgCode(args Args) interface{} {

	model := models.TgLotteryOrgCode{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindLotteryPart(args Args) interface{} {

	model := models.TgLotteryPart{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindLotteryPrize(args Args) interface{} {

	model := models.TgLotteryPrize{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindLotteryRecord(args Args) interface{} {

	model := models.TgLotteryRecord{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindLotteryThread(args Args) interface{} {

	model := models.TgLotteryThread{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindLotteryUser(args Args) interface{} {

	model := models.TgLotteryUser{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindLotteryUserCode(args Args) interface{} {

	model := models.TgLotteryUserCode{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindLoveChicken(args Args) interface{} {

	model := models.TgLoveChicken{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindLoverUser(args Args) interface{} {

	model := models.TgLoverUser{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindMainWork(args Args) interface{} {

	model := models.TgMainWork{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindManagementType(args Args) interface{} {

	model := models.TgManagementType{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindManualCode(args Args) interface{} {

	model := models.TgManualCode{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindManualCodeOperationRecord(args Args) interface{} {

	model := models.TgManualCodeOperationRecord{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindMenu(args Args) interface{} {

	model := models.TgMenu{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindMenuCopy(args Args) interface{} {

	model := models.TgMenuCopy{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindMessage(args Args) interface{} {

	model := models.TgMessage{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindMessageApp(args Args) interface{} {

	model := models.TgMessageApp{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindMessageEvent(args Args) interface{} {

	model := models.TgMessageEvent{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindMessageEventUser(args Args) interface{} {

	model := models.TgMessageEventUser{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindMessageExtend(args Args) interface{} {

	model := models.TgMessageExtend{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindMessageImg(args Args) interface{} {

	model := models.TgMessageImg{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindMessageInteract(args Args) interface{} {

	model := models.TgMessageInteract{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindMessageInteractUser(args Args) interface{} {

	model := models.TgMessageInteractUser{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindMessageJike(args Args) interface{} {

	model := models.TgMessageJike{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindMessageJikeUser(args Args) interface{} {

	model := models.TgMessageJikeUser{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindMessageLabel(args Args) interface{} {

	model := models.TgMessageLabel{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindMessageLovecar(args Args) interface{} {

	model := models.TgMessageLovecar{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindMessageLovecarUser(args Args) interface{} {

	model := models.TgMessageLovecarUser{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindMessageRecord(args Args) interface{} {

	model := models.TgMessageRecord{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindMessageSetting(args Args) interface{} {

	model := models.TgMessageSetting{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindMessageStore(args Args) interface{} {

	model := models.TgMessageStore{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindMessageStoreUser(args Args) interface{} {

	model := models.TgMessageStoreUser{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindMessageType(args Args) interface{} {

	model := models.TgMessageType{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindMessageTypeContent(args Args) interface{} {

	model := models.TgMessageTypeContent{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindMessageTypeCopy(args Args) interface{} {

	model := models.TgMessageTypeCopy{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindMessageUrl(args Args) interface{} {

	model := models.TgMessageUrl{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindMsgBroker(args Args) interface{} {

	model := models.TgMsgBroker{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindMusic(args Args) interface{} {

	model := models.TgMusic{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindMusic1(args Args) interface{} {

	model := models.TgMusic1{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindMusicType(args Args) interface{} {

	model := models.TgMusicType{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindNodePunchAuth(args Args) interface{} {

	model := models.TgNodePunchAuth{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindNodePunchConfirm(args Args) interface{} {

	model := models.TgNodePunchConfirm{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindOnlineserviceMute(args Args) interface{} {

	model := models.TgOnlineserviceMute{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindOnlineserviceOrder(args Args) interface{} {

	model := models.TgOnlineserviceOrder{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindOperateActivity(args Args) interface{} {

	model := models.TgOperateActivity{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindOperateActivityCheck(args Args) interface{} {

	model := models.TgOperateActivityCheck{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindOrder(args Args) interface{} {

	model := models.TgOrder{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindOrderMain(args Args) interface{} {

	model := models.TgOrderMain{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindOrgBankCard(args Args) interface{} {

	model := models.TgOrgBankCard{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindOrgMoney(args Args) interface{} {

	model := models.TgOrgMoney{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindOrgMoneyBill(args Args) interface{} {

	model := models.TgOrgMoneyBill{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindOrgMoneyRecharge(args Args) interface{} {

	model := models.TgOrgMoneyRecharge{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindOrgMoneyWithdraw(args Args) interface{} {

	model := models.TgOrgMoneyWithdraw{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindOrganization(args Args) interface{} {

	model := models.TgOrganization{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindOrganizationActivity(args Args) interface{} {

	model := models.TgOrganizationActivity{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindOrganizationCar(args Args) interface{} {

	model := models.TgOrganizationCar{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindOrganizationCheck(args Args) interface{} {

	model := models.TgOrganizationCheck{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindOrganizationCheckInfo(args Args) interface{} {

	model := models.TgOrganizationCheckInfo{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindOrganizationConf(args Args) interface{} {

	model := models.TgOrganizationConf{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindOrganizationConfVal(args Args) interface{} {

	model := models.TgOrganizationConfVal{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindOrganizationCopy(args Args) interface{} {

	model := models.TgOrganizationCopy{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindOrganizationCopyCopy(args Args) interface{} {

	model := models.TgOrganizationCopyCopy{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindOrganizationFocus(args Args) interface{} {

	model := models.TgOrganizationFocus{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindOrganizationLog(args Args) interface{} {

	model := models.TgOrganizationLog{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindOrganizationOrder(args Args) interface{} {

	model := models.TgOrganizationOrder{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindOrganizationOrderpointsRecord(args Args) interface{} {

	model := models.TgOrganizationOrderpointsRecord{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindOrganizationPrivilege(args Args) interface{} {

	model := models.TgOrganizationPrivilege{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindOrganizationPrivilegeShop(args Args) interface{} {

	model := models.TgOrganizationPrivilegeShop{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindOrganizationPrivilegeStore(args Args) interface{} {

	model := models.TgOrganizationPrivilegeStore{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindOrganizationTime(args Args) interface{} {

	model := models.TgOrganizationTime{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindOrganizationType(args Args) interface{} {

	model := models.TgOrganizationType{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindPartner(args Args) interface{} {

	model := models.TgPartner{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindPartnerBlacklist(args Args) interface{} {

	model := models.TgPartnerBlacklist{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindPartnerInvite(args Args) interface{} {

	model := models.TgPartnerInvite{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindPartnerLog(args Args) interface{} {

	model := models.TgPartnerLog{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindPartnerPeriod(args Args) interface{} {

	model := models.TgPartnerPeriod{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindPartnerQrcode(args Args) interface{} {

	model := models.TgPartnerQrcode{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindPartnerRegister(args Args) interface{} {

	model := models.TgPartnerRegister{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindPartnerResults(args Args) interface{} {

	model := models.TgPartnerResults{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindPartnerTask(args Args) interface{} {

	model := models.TgPartnerTask{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindPayAlipayLog(args Args) interface{} {

	model := models.TgPayAlipayLog{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindPayWxLog(args Args) interface{} {

	model := models.TgPayWxLog{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindPhone(args Args) interface{} {

	model := models.TgPhone{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindPhoto(args Args) interface{} {

	model := models.TgPhoto{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindPointsData(args Args) interface{} {

	model := models.TgPointsData{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindPointsExceed(args Args) interface{} {

	model := models.TgPointsExceed{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindPointsExpRecords(args Args) interface{} {

	model := models.TgPointsExpRecords{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindPointsExpRules(args Args) interface{} {

	model := models.TgPointsExpRules{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindPointsGoods(args Args) interface{} {

	model := models.TgPointsGoods{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindPointsGoodsPay(args Args) interface{} {

	model := models.TgPointsGoodsPay{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindPointsLog(args Args) interface{} {

	model := models.TgPointsLog{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindPointsProduct(args Args) interface{} {

	model := models.TgPointsProduct{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindPointsProductRecords(args Args) interface{} {

	model := models.TgPointsProductRecords{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindPointsRanking(args Args) interface{} {

	model := models.TgPointsRanking{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindPointsSaleRules(args Args) interface{} {

	model := models.TgPointsSaleRules{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindPraise(args Args) interface{} {

	model := models.TgPraise{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindProjectExamine(args Args) interface{} {

	model := models.TgProjectExamine{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindPunch(args Args) interface{} {

	model := models.TgPunch{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindPurchase(args Args) interface{} {

	model := models.TgPurchase{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindPurchaseFactory(args Args) interface{} {

	model := models.TgPurchaseFactory{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindPurchaseFail(args Args) interface{} {

	model := models.TgPurchaseFail{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindPurchaseImg(args Args) interface{} {

	model := models.TgPurchaseImg{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindPurchaseRecords(args Args) interface{} {

	model := models.TgPurchaseRecords{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindPurchaseSms(args Args) interface{} {

	model := models.TgPurchaseSms{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindPurchaseUser(args Args) interface{} {

	model := models.TgPurchaseUser{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindPushSmsConfig(args Args) interface{} {

	model := models.TgPushSmsConfig{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindQuestionnaire(args Args) interface{} {

	model := models.TgQuestionnaire{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindQuestionnaireAnswer(args Args) interface{} {

	model := models.TgQuestionnaireAnswer{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindQuestionnaireMain(args Args) interface{} {

	model := models.TgQuestionnaireMain{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindQuestionnaireUser(args Args) interface{} {

	model := models.TgQuestionnaireUser{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindRabbitMsg(args Args) interface{} {

	model := models.TgRabbitMsg{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindRecommend(args Args) interface{} {

	model := models.TgRecommend{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindRecommendAuth(args Args) interface{} {

	model := models.TgRecommendAuth{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindRecommendContent(args Args) interface{} {

	model := models.TgRecommendContent{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindRecommendDetail201905(args Args) interface{} {

	model := models.TgRecommendDetail201905{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindRecommendDetail201906(args Args) interface{} {

	model := models.TgRecommendDetail201906{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindRecommendDetail201907(args Args) interface{} {

	model := models.TgRecommendDetail201907{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindRecommendDetail201908(args Args) interface{} {

	model := models.TgRecommendDetail201908{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindRecommendDetail201909(args Args) interface{} {

	model := models.TgRecommendDetail201909{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindRecommendDetail201910(args Args) interface{} {

	model := models.TgRecommendDetail201910{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindRecommendDetail201911(args Args) interface{} {

	model := models.TgRecommendDetail201911{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindRecommendDetail201912(args Args) interface{} {

	model := models.TgRecommendDetail201912{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindRecommendOrg(args Args) interface{} {

	model := models.TgRecommendOrg{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindRecommendPosition(args Args) interface{} {

	model := models.TgRecommendPosition{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindRecommendPositionConfig(args Args) interface{} {

	model := models.TgRecommendPositionConfig{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindRecommendRecord(args Args) interface{} {

	model := models.TgRecommendRecord{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindRecommendTime(args Args) interface{} {

	model := models.TgRecommendTime{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindRecommendTimeExtend(args Args) interface{} {

	model := models.TgRecommendTimeExtend{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindRecommendTongjiDay(args Args) interface{} {

	model := models.TgRecommendTongjiDay{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindRecommendTongjiMonth(args Args) interface{} {

	model := models.TgRecommendTongjiMonth{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindRedpack(args Args) interface{} {

	model := models.TgRedpack{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindRedpackArea(args Args) interface{} {

	model := models.TgRedpackArea{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindRedpackCheckRecord(args Args) interface{} {

	model := models.TgRedpackCheckRecord{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindRedpackContent(args Args) interface{} {

	model := models.TgRedpackContent{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindRedpackMsgRecord(args Args) interface{} {

	model := models.TgRedpackMsgRecord{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindRedpackOrg(args Args) interface{} {

	model := models.TgRedpackOrg{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindRedpackPosition(args Args) interface{} {

	model := models.TgRedpackPosition{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindRedpackPrize(args Args) interface{} {

	model := models.TgRedpackPrize{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindRedpackRecord(args Args) interface{} {

	model := models.TgRedpackRecord{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindRedpackRecordTemp(args Args) interface{} {

	model := models.TgRedpackRecordTemp{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindRedpackTime(args Args) interface{} {

	model := models.TgRedpackTime{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindRedpackUserCard(args Args) interface{} {

	model := models.TgRedpackUserCard{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindRedpackUserCardRecord(args Args) interface{} {

	model := models.TgRedpackUserCardRecord{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindRegion(args Args) interface{} {

	model := models.TgRegion{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindReport(args Args) interface{} {

	model := models.TgReport{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindReportImg(args Args) interface{} {

	model := models.TgReportImg{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindReportRecord(args Args) interface{} {

	model := models.TgReportRecord{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindRide(args Args) interface{} {

	model := models.TgRide{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindRideCheck(args Args) interface{} {

	model := models.TgRideCheck{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindRideCheckNotes(args Args) interface{} {

	model := models.TgRideCheckNotes{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindRideLine(args Args) interface{} {

	model := models.TgRideLine{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindRideNode(args Args) interface{} {

	model := models.TgRideNode{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindRoadBook(args Args) interface{} {

	model := models.TgRoadBook{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindRoadBookLine(args Args) interface{} {

	model := models.TgRoadBookLine{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindRoadBookLineContent(args Args) interface{} {

	model := models.TgRoadBookLineContent{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindRoadBookMeet(args Args) interface{} {

	model := models.TgRoadBookMeet{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindRoadBookRecommend(args Args) interface{} {

	model := models.TgRoadBookRecommend{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindRoadbookArouse(args Args) interface{} {

	model := models.TgRoadbookArouse{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindRoadbookArouseAbnormal(args Args) interface{} {

	model := models.TgRoadbookArouseAbnormal{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindRoadbookArouseAccount(args Args) interface{} {

	model := models.TgRoadbookArouseAccount{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindRoadbookArouseAttend(args Args) interface{} {

	model := models.TgRoadbookArouseAttend{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindRoadbookArouseAttendCensus(args Args) interface{} {

	model := models.TgRoadbookArouseAttendCensus{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindRoadbookArouseFlow(args Args) interface{} {

	model := models.TgRoadbookArouseFlow{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindRoadbookArouseLog(args Args) interface{} {

	model := models.TgRoadbookArouseLog{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindRoadbookArouseRealAccount(args Args) interface{} {

	model := models.TgRoadbookArouseRealAccount{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindRoadbookArouseRelated(args Args) interface{} {

	model := models.TgRoadbookArouseRelated{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindRoadbookLine(args Args) interface{} {

	model := models.TgRoadbookLine{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindRoadbookMerchant(args Args) interface{} {

	model := models.TgRoadbookMerchant{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindRoadbookMerchantApply(args Args) interface{} {

	model := models.TgRoadbookMerchantApply{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindRoadbookMerchantInvite(args Args) interface{} {

	model := models.TgRoadbookMerchantInvite{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindRoadbookNodePunchCensus(args Args) interface{} {

	model := models.TgRoadbookNodePunchCensus{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindRouletteLottery(args Args) interface{} {

	model := models.TgRouletteLottery{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindRouletteLotteryAttend(args Args) interface{} {

	model := models.TgRouletteLotteryAttend{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindRouletteLotteryAward(args Args) interface{} {

	model := models.TgRouletteLotteryAward{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindSaleinfoOrg(args Args) interface{} {

	model := models.TgSaleinfoOrg{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindSalerStar(args Args) interface{} {

	model := models.TgSalerStar{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindSearchLog(args Args) interface{} {

	model := models.TgSearchLog{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindSeries(args Args) interface{} {

	model := models.TgSeries{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindServiceCity(args Args) interface{} {

	model := models.TgServiceCity{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindSmsApiLog(args Args) interface{} {

	model := models.TgSmsApiLog{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindSmsTemplate(args Args) interface{} {

	model := models.TgSmsTemplate{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindSmsTemplateLog(args Args) interface{} {

	model := models.TgSmsTemplateLog{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindSmsTotal(args Args) interface{} {

	model := models.TgSmsTotal{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindSmsVerifyCode(args Args) interface{} {

	model := models.TgSmsVerifyCode{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindSpecial(args Args) interface{} {

	model := models.TgSpecial{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindSpmAppDownloadLog(args Args) interface{} {

	model := models.TgSpmAppDownloadLog{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindSpmBaidurefererLog(args Args) interface{} {

	model := models.TgSpmBaidurefererLog{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindSpmCate(args Args) interface{} {

	model := models.TgSpmCate{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindSpmDesignation(args Args) interface{} {

	model := models.TgSpmDesignation{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindSpmKeywords(args Args) interface{} {

	model := models.TgSpmKeywords{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindSpmLog(args Args) interface{} {

	model := models.TgSpmLog{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindSpmLogCopy(args Args) interface{} {

	model := models.TgSpmLogCopy{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindSpmParameter(args Args) interface{} {

	model := models.TgSpmParameter{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindSpmPlace(args Args) interface{} {

	model := models.TgSpmPlace{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindSpmPromoter(args Args) interface{} {

	model := models.TgSpmPromoter{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindStaff(args Args) interface{} {

	model := models.TgStaff{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindStatistics(args Args) interface{} {

	model := models.TgStatistics{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindSupertopicApply(args Args) interface{} {

	model := models.TgSupertopicApply{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindTagCar(args Args) interface{} {

	model := models.TgTagCar{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindTagCartSpecific(args Args) interface{} {

	model := models.TgTagCartSpecific{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindTags(args Args) interface{} {

	model := models.TgTags{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindTagsCategory(args Args) interface{} {

	model := models.TgTagsCategory{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindTagscategory(args Args) interface{} {

	model := models.TgTagscategory{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindThreadActivity(args Args) interface{} {

	model := models.TgThreadActivity{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindThreadActivityApply(args Args) interface{} {

	model := models.TgThreadActivityApply{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindThreadActivityCar(args Args) interface{} {

	model := models.TgThreadActivityCar{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindThreadActivityMaster(args Args) interface{} {

	model := models.TgThreadActivityMaster{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindThreadActivityScene(args Args) interface{} {

	model := models.TgThreadActivityScene{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindThreadActivityType(args Args) interface{} {

	model := models.TgThreadActivityType{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindThreadActivityUser(args Args) interface{} {

	model := models.TgThreadActivityUser{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindThreadAsk(args Args) interface{} {

	model := models.TgThreadAsk{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindThreadAskAnswer(args Args) interface{} {

	model := models.TgThreadAskAnswer{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindThreadAskAnswerFollow(args Args) interface{} {

	model := models.TgThreadAskAnswerFollow{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindThreadCensusRanking(args Args) interface{} {

	model := models.TgThreadCensusRanking{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindThreadCycle(args Args) interface{} {

	model := models.TgThreadCycle{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindThreadDeleteReason(args Args) interface{} {

	model := models.TgThreadDeleteReason{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindThreadEvaluate(args Args) interface{} {

	model := models.TgThreadEvaluate{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindThreadEvaluateCar(args Args) interface{} {

	model := models.TgThreadEvaluateCar{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindThreadGoods(args Args) interface{} {

	model := models.TgThreadGoods{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindThreadGoodsCategory(args Args) interface{} {

	model := models.TgThreadGoodsCategory{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindThreadIntegralLog(args Args) interface{} {

	model := models.TgThreadIntegralLog{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindThreadMastersay(args Args) interface{} {

	model := models.TgThreadMastersay{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindThreadMastersayCar(args Args) interface{} {

	model := models.TgThreadMastersayCar{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindThreadMood(args Args) interface{} {

	model := models.TgThreadMood{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindThreadPlay(args Args) interface{} {

	model := models.TgThreadPlay{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}


func BindThreadRoadbookAnnex(args Args) interface{} {

	model := models.TgThreadRoadbookAnnex{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindThreadRoadbookLineBeen(args Args) interface{} {

	model := models.TgThreadRoadbookLineBeen{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindThreadRoadbookLineRelated(args Args) interface{} {

	model := models.TgThreadRoadbookLineRelated{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindThreadRoadbookNoteUser(args Args) interface{} {

	model := models.TgThreadRoadbookNoteUser{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindThreadRoadbookPunchCensus(args Args) interface{} {

	model := models.TgThreadRoadbookPunchCensus{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindThreadShow(args Args) interface{} {

	model := models.TgThreadShow{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindThreadTopic(args Args) interface{} {

	model := models.TgThreadTopic{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindThreadTopicAccess(args Args) interface{} {

	model := models.TgThreadTopicAccess{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindThreadTopicApply(args Args) interface{} {

	model := models.TgThreadTopicApply{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindThreadTopicFans(args Args) interface{} {

	model := models.TgThreadTopicFans{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindThreadVisible(args Args) interface{} {

	model := models.TgThreadVisible{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindThreadWordColumn(args Args) interface{} {

	model := models.TgThreadWordColumn{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindThreadWords(args Args) interface{} {

	model := models.TgThreadWords{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindTicket(args Args) interface{} {

	model := models.TgTicket{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindTicketApiLog(args Args) interface{} {

	model := models.TgTicketApiLog{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindTicketApplyUser(args Args) interface{} {

	model := models.TgTicketApplyUser{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindTicketCheckRecord(args Args) interface{} {

	model := models.TgTicketCheckRecord{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindTicketCheckoutAuth(args Args) interface{} {

	model := models.TgTicketCheckoutAuth{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindTicketFree(args Args) interface{} {

	model := models.TgTicketFree{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindTicketFreeHandout(args Args) interface{} {

	model := models.TgTicketFreeHandout{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindTicketNumsRecord(args Args) interface{} {

	model := models.TgTicketNumsRecord{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindTicketPaperRecord(args Args) interface{} {

	model := models.TgTicketPaperRecord{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindTicketQd(args Args) interface{} {

	model := models.TgTicketQd{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindTicketQdSale(args Args) interface{} {

	model := models.TgTicketQdSale{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindTicketRedpack(args Args) interface{} {

	model := models.TgTicketRedpack{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindTicketRedpackLog(args Args) interface{} {

	model := models.TgTicketRedpackLog{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindTicketRefund(args Args) interface{} {

	model := models.TgTicketRefund{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindTicketRefundLog(args Args) interface{} {

	model := models.TgTicketRefundLog{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindTicketSendRecord(args Args) interface{} {

	model := models.TgTicketSendRecord{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindTicketThirdParty(args Args) interface{} {

	model := models.TgTicketThirdParty{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindTicketUser(args Args) interface{} {

	model := models.TgTicketUser{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindTmpData(args Args) interface{} {

	model := models.TgTmpData{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindTopicCategory(args Args) interface{} {

	model := models.TgTopicCategory{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindTopicExpLog(args Args) interface{} {

	model := models.TgTopicExpLog{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindTopicLog(args Args) interface{} {

	model := models.TgTopicLog{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindTopicScompere(args Args) interface{} {

	model := models.TgTopicScompere{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindTopicSignin(args Args) interface{} {

	model := models.TgTopicSignin{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindTrafficStatistics(args Args) interface{} {

	model := models.TgTrafficStatistics{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindTrafficStatisticsLog(args Args) interface{} {

	model := models.TgTrafficStatisticsLog{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindTrafficStatisticsLogGroup(args Args) interface{} {

	model := models.TgTrafficStatisticsLogGroup{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindTwitter(args Args) interface{} {

	model := models.TgTwitter{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindTwitterAlbum(args Args) interface{} {

	model := models.TgTwitterAlbum{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindTwitterSource(args Args) interface{} {

	model := models.TgTwitterSource{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindTwitterTopicAccess(args Args) interface{} {

	model := models.TgTwitterTopicAccess{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindTwitterVisibility(args Args) interface{} {

	model := models.TgTwitterVisibility{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindUploadApk(args Args) interface{} {

	model := models.TgUploadApk{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindUserBankCard(args Args) interface{} {

	model := models.TgUserBankCard{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindUserBehaviour(args Args) interface{} {

	model := models.TgUserBehaviour{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindUserCar(args Args) interface{} {

	model := models.TgUserCar{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindUserFollow(args Args) interface{} {

	model := models.TgUserFollow{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindUserPoints(args Args) interface{} {

	model := models.TgUserPoints{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindUserRecommend(args Args) interface{} {

	model := models.TgUserRecommend{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindUserSetting(args Args) interface{} {

	model := models.TgUserSetting{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindVideo(args Args) interface{} {

	model := models.TgVideo{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindVideoAttach(args Args) interface{} {

	model := models.TgVideoAttach{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindVideoAttachReceive(args Args) interface{} {

	model := models.TgVideoAttachReceive{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindVideoLog(args Args) interface{} {

	model := models.TgVideoLog{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindVideoNotice(args Args) interface{} {

	model := models.TgVideoNotice{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindVideoRecommendTable(args Args) interface{} {

	model := models.TgVideoRecommendTable{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindVideoVisible(args Args) interface{} {

	model := models.TgVideoVisible{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindVirtualNickname(args Args) interface{} {

	model := models.TgVirtualNickname{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindVirtualRecord(args Args) interface{} {

	model := models.TgVirtualRecord{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindWannaBuy(args Args) interface{} {

	model := models.TgWannaBuy{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindWeatherCitys(args Args) interface{} {

	model := models.TgWeatherCitys{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindWeatherToday(args Args) interface{} {

	model := models.TgWeatherToday{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindWeixinConfig(args Args) interface{} {

	model := models.TgWeixinConfig{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindWeixinCustomSend(args Args) interface{} {

	model := models.TgWeixinCustomSend{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindWeixinKerword(args Args) interface{} {

	model := models.TgWeixinKerword{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindWeixinMaterial(args Args) interface{} {

	model := models.TgWeixinMaterial{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindWeixinMenu(args Args) interface{} {

	model := models.TgWeixinMenu{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindWeixinPush(args Args) interface{} {

	model := models.TgWeixinPush{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindWeixinPushAuto(args Args) interface{} {

	model := models.TgWeixinPushAuto{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindWeixinPushConfig(args Args) interface{} {

	model := models.TgWeixinPushConfig{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindWeixinPushInteractive(args Args) interface{} {

	model := models.TgWeixinPushInteractive{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindWeixinPushMaterial(args Args) interface{} {

	model := models.TgWeixinPushMaterial{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindWeixinPushRecord(args Args) interface{} {

	model := models.TgWeixinPushRecord{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindWeixinPushTag(args Args) interface{} {

	model := models.TgWeixinPushTag{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindWeixinQrcode(args Args) interface{} {

	model := models.TgWeixinQrcode{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindWeixinQrcodeScanLog(args Args) interface{} {

	model := models.TgWeixinQrcodeScanLog{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindWeixinTemplate(args Args) interface{} {

	model := models.TgWeixinTemplate{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindXc1Jurisdiction(args Args) interface{} {

	model := models.TgXc1Jurisdiction{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindXc1ManageObject(args Args) interface{} {

	model := models.TgXc1ManageObject{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindXc1Org2scene(args Args) interface{} {

	model := models.TgXc1Org2scene{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindXc1Worker2guest(args Args) interface{} {

	model := models.TgXc1Worker2guest{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindXcAppMessage(args Args) interface{} {

	model := models.TgXcAppMessage{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindXcAppMessageCount(args Args) interface{} {

	model := models.TgXcAppMessageCount{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindXcApplyItem(args Args) interface{} {

	model := models.TgXcApplyItem{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindXcApplyType(args Args) interface{} {

	model := models.TgXcApplyType{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindXcArea(args Args) interface{} {

	model := models.TgXcArea{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindXcAuthMapping(args Args) interface{} {

	model := models.TgXcAuthMapping{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindXcAuthorise(args Args) interface{} {

	model := models.TgXcAuthorise{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindXcCheckCourse(args Args) interface{} {

	model := models.TgXcCheckCourse{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindXcCheckTable(args Args) interface{} {

	model := models.TgXcCheckTable{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindXcChecker2object(args Args) interface{} {

	model := models.TgXcChecker2object{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindXcFriends(args Args) interface{} {

	model := models.TgXcFriends{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindXcGuest2object(args Args) interface{} {

	model := models.TgXcGuest2object{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindXcJobType(args Args) interface{} {

	model := models.TgXcJobType{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindXcMessage(args Args) interface{} {

	model := models.TgXcMessage{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindXcMessageAddtask(args Args) interface{} {

	model := models.TgXcMessageAddtask{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindXcMessageApply(args Args) interface{} {

	model := models.TgXcMessageApply{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindXcMessageContent(args Args) interface{} {

	model := models.TgXcMessageContent{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindXcMessageRectification(args Args) interface{} {

	model := models.TgXcMessageRectification{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindXcMessageSetting(args Args) interface{} {

	model := models.TgXcMessageSetting{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindXcMessageUnqualifiedcheck(args Args) interface{} {

	model := models.TgXcMessageUnqualifiedcheck{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindXcObject(args Args) interface{} {

	model := models.TgXcObject{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindXcObject2dispose(args Args) interface{} {

	model := models.TgXcObject2dispose{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindXcObjectOrganization(args Args) interface{} {

	model := models.TgXcObjectOrganization{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindXcObjectType(args Args) interface{} {

	model := models.TgXcObjectType{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindXcOffsetTriggerTime(args Args) interface{} {

	model := models.TgXcOffsetTriggerTime{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindXcOrder(args Args) interface{} {

	model := models.TgXcOrder{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindXcOrderProgress(args Args) interface{} {

	model := models.TgXcOrderProgress{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindXcOrganization(args Args) interface{} {

	model := models.TgXcOrganization{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindXcPayLog(args Args) interface{} {

	model := models.TgXcPayLog{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindXcPermission(args Args) interface{} {

	model := models.TgXcPermission{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindXcScene(args Args) interface{} {

	model := models.TgXcScene{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindXcSceneReview(args Args) interface{} {

	model := models.TgXcSceneReview{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindXcStaff2object(args Args) interface{} {

	model := models.TgXcStaff2object{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindXcTask(args Args) interface{} {

	model := models.TgXcTask{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindXcTaskCopy(args Args) interface{} {

	model := models.TgXcTaskCopy{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindXcTaskDis(args Args) interface{} {

	model := models.TgXcTaskDis{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindXcTaskDisimg(args Args) interface{} {

	model := models.TgXcTaskDisimg{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindXcTaskLog(args Args) interface{} {

	model := models.TgXcTaskLog{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindXcTaskPackage(args Args) interface{} {

	model := models.TgXcTaskPackage{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindXcTopUpLog(args Args) interface{} {

	model := models.TgXcTopUpLog{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindXcUser2scene(args Args) interface{} {

	model := models.TgXcUser2scene{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindXcUsercoordinate(args Args) interface{} {

	model := models.TgXcUsercoordinate{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindXcWorker(args Args) interface{} {

	model := models.TgXcWorker{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindXcWorker2area(args Args) interface{} {

	model := models.TgXcWorker2area{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindXcWorkerYanghuanBack(args Args) interface{} {

	model := models.TgXcWorkerYanghuanBack{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindYaoyiyao(args Args) interface{} {

	model := models.TgYaoyiyao{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindYaoyiyaoAuth(args Args) interface{} {

	model := models.TgYaoyiyaoAuth{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindYaoyiyaoAwards(args Args) interface{} {

	model := models.TgYaoyiyaoAwards{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindYaoyiyaoAwardsCopy(args Args) interface{} {

	model := models.TgYaoyiyaoAwardsCopy{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindYaoyiyaoCacheawards(args Args) interface{} {

	model := models.TgYaoyiyaoCacheawards{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindYaoyiyaoEndLog(args Args) interface{} {

	model := models.TgYaoyiyaoEndLog{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindYaoyiyaoNotWin(args Args) interface{} {

	model := models.TgYaoyiyaoNotWin{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindYaoyiyaoPassLog(args Args) interface{} {

	model := models.TgYaoyiyaoPassLog{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindYaoyiyaoRank(args Args) interface{} {

	model := models.TgYaoyiyaoRank{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindYaoyiyaoRotary(args Args) interface{} {

	model := models.TgYaoyiyaoRotary{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindYaoyiyaoSignin(args Args) interface{} {

	model := models.TgYaoyiyaoSignin{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}

func BindYaoyiyaoStatus(args Args) interface{} {

	model := models.TgYaoyiyaoStatus{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}