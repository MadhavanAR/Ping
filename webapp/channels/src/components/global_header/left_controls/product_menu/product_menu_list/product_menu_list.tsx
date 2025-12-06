// Copyright (c) 2015-present Mattermost, Inc. All Rights Reserved.
// See LICENSE.txt for license information.

import React, {useEffect} from 'react';
import {useIntl} from 'react-intl';
import {useSelector} from 'react-redux';

import {
    AccountMultipleOutlineIcon,
    ApplicationCogIcon,
    DownloadOutlineIcon,
    InformationOutlineIcon,
    ViewGridPlusOutlineIcon,
    WebhookIncomingIcon,
} from '@mattermost/compass-icons/components';
import type {UserProfile} from '@mattermost/types/users';

import {Permissions} from 'mattermost-redux/constants';
import {isCurrentUserSystemAdmin} from 'mattermost-redux/selectors/entities/users';

import AboutBuildModal from 'components/about_build_modal';
import {VisitSystemConsoleTour} from 'components/onboarding_tasks';
import SystemPermissionGate from 'components/permissions_gates/system_permission_gate';
import TeamPermissionGate from 'components/permissions_gates/team_permission_gate';
import UserGroupsModal from 'components/user_groups_modal';
import Menu from 'components/widgets/menu/menu';
import RestrictedIndicator from 'components/widgets/menu/menu_items/restricted_indicator';

import {FREEMIUM_TO_ENTERPRISE_TRIAL_LENGTH_DAYS} from 'utils/cloud_utils';
import {LicenseSkus, ModalIdentifiers, MattermostFeatures} from 'utils/constants';
import {makeUrlSafe} from 'utils/url';
import * as UserAgent from 'utils/user_agent';

import type {ModalData} from 'types/actions';

import './product_menu_list.scss';

export type Props = {
    isMobile: boolean;
    teamId?: string;
    teamName?: string;
    siteName: string;
    currentUser: UserProfile;
    appDownloadLink: string;
    isMessaging: boolean;
    enableCommands: boolean;
    enableIncomingWebhooks: boolean;
    enableOAuthServiceProvider: boolean;
    enableOutgoingWebhooks: boolean;
    canManageSystemBots: boolean;
    canManageIntegrations: boolean;
    enablePluginMarketplace: boolean;
    showVisitSystemConsoleTour: boolean;
    isStarterFree: boolean;
    isFreeTrial: boolean;
    onClick?: React.MouseEventHandler<HTMLElement>;
    handleVisitConsoleClick: React.MouseEventHandler<HTMLElement>;
    enableCustomUserGroups?: boolean;
    actions: {
        openModal: <P>(modalData: ModalData<P>) => void;
        getPrevTrialLicense: () => void;
    };
};

const ProductMenuList = (props: Props): JSX.Element | null => {
    const {
        teamId,
        teamName,
        siteName,
        currentUser,
        appDownloadLink,
        isMessaging,
        enableCommands,
        enableIncomingWebhooks,
        enableOAuthServiceProvider,
        enableOutgoingWebhooks,
        canManageSystemBots,
        canManageIntegrations,
        enablePluginMarketplace,
        showVisitSystemConsoleTour,
        isStarterFree,
        isFreeTrial,
        onClick,
        handleVisitConsoleClick,
        isMobile = false,
        enableCustomUserGroups,
    } = props;
    const {formatMessage} = useIntl();
    const isAdmin = useSelector(isCurrentUserSystemAdmin);

    useEffect(() => {
        props.actions.getPrevTrialLicense();
    }, []);

    if (!currentUser) {
        return null;
    }

    const openGroupsModal = () => {
        props.actions.openModal({
            modalId: ModalIdentifiers.USER_GROUPS,
            dialogType: UserGroupsModal,
            dialogProps: {
                backButtonAction: openGroupsModal,
            },
        });
    };

    const someIntegrationEnabled = enableIncomingWebhooks || enableOutgoingWebhooks || enableCommands || enableOAuthServiceProvider || canManageSystemBots;
    const showIntegrations = !isMobile && someIntegrationEnabled && canManageIntegrations;

    return (
        <Menu.Group>
            <div onClick={onClick}>
                {/* Removed: CloudTrial, ItemCloudLimit, Integrations, Marketplace, UserGroups with trial restrictions */}
                <SystemPermissionGate permissions={Permissions.SYSCONSOLE_READ_PERMISSIONS}>
                    <Menu.ItemLink
                        id='systemConsole'
                        show={!isMobile}
                        to='/admin_console'
                        text={(
                            <>
                                {formatMessage({id: 'navbar_dropdown.console', defaultMessage: 'System Console'})}
                                {showVisitSystemConsoleTour && (
                                    <div
                                        onClick={handleVisitConsoleClick}
                                        className={'system-console-visit'}
                                    >
                                        <VisitSystemConsoleTour/>
                                    </div>
                                )}
                            </>
                        )}
                        icon={<ApplicationCogIcon size={18}/>}
                    />
                </SystemPermissionGate>
                <Menu.ItemExternalLink
                    id='nativeAppLink'
                    show={appDownloadLink && !UserAgent.isMobileApp()}
                    url={makeUrlSafe(appDownloadLink)}
                    text={formatMessage({id: 'navbar_dropdown.nativeApps', defaultMessage: 'Download Apps'})}
                    icon={<DownloadOutlineIcon size={18}/>}
                />
                <Menu.ItemToggleModalRedux
                    id='about'
                    modalId={ModalIdentifiers.ABOUT}
                    dialogType={AboutBuildModal}
                    text={formatMessage({id: 'navbar_dropdown.about', defaultMessage: 'About {appTitle}'}, {appTitle: siteName})}
                    icon={<InformationOutlineIcon size={18}/>}
                />
            </div>
        </Menu.Group>
    );
};

export default ProductMenuList;
