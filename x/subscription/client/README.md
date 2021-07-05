# Subscription Module
This document introduces the **Subscription** module of the Sentinel Hub. Note that all examples use the official sentinel rpc server (https://rpc.sentinel.co:443). Obviously, any other rpc server can be used. Note further that all example transactions are signed with a local key named *myKeyringWallet*. Replace this with your own key name. Keys can be managed using the *keys* module. At the time of writing, the chain-id *sentinelhub-2* is used. This may change in the future.

		
# Queries
The **Queries** submodule provides the logic to request information from subscriptions and quotas. The following messages are available:
- [Subscription](#subscription)
- [Subscriptions](#subscriptions)
- [Quota](#quota)
- [Quotas](#quotas)


## Subscription

**Description**

Query a single subscription.

**Usage**	

	sentinelcli query subscription <SUBSCRIPTION_ID> [FLAGS]

**Example**

	// Query subscription 1.
	sentinelcli query subscription 1 --node https://rpc.sentinel.co:443


## Subscriptions

**Description**

Query multiple subscriptions.

**Usage**

	sentinelcli query subscriptions [FLAGS]

**Flags**

- **status** 		[STATUS]		Filter result based on subscription status
- **plan**		[PLAN_ID]		Filter result based on plan
- **address**		[ACCOUNT_ADDRESS]	Filter result based on owner address
- **page**		[PAGE_NUMBER]		Show results of a specific page number

**Example**

	// Query all subscriptions.
	sentinelcli query subscriptions --node https://rpc.sentinel.co:443


## Quota

**Description**

Query a single quota of a subscribed address.

**Usage**

	sentinelcli query quota <SUBSCRIPTION_ID> <ACCOUNT_ADDRESS> [FLAGS]

**Example**

	// Query the quota of subscription 1 and address sent1v5ktkl3x9z2ky5uxhqjy487t7x0gl7fjrm0npv.
	sentinelcli query quota 1 sent1v5ktkl3x9z2ky5uxhqjy487t7x0gl7fjrm0npv --node https://rpc.sentinel.co:443


## Quotas

**Description**

Query all quotas of a subscription.

**Usage**

	sentinelcli query quotas <SUBSCRIPTION_ID> [FLAGS]

**Flags**

- **page**		[PAGE_NUMBER]		Show results of a specific page number

**Example**
	
	// Query the quotas of subscription 1.
	sentinelcli query quotas 1 --node https://rpc.sentinel.co:443


# Transactions
The **Transactions** submodule provides the logic to manage subscriptions and quotas. The following messages are available:
- [Quota Add](#quota-add)
- [Quota Update](#quota-update)
- [Subscribe To Node](#subscribe-to-node)
- [Subscribe to Plan](#subscribe-to-plan)
- [Cancel](#cancel) 


## Quota Add

**Description**

Add a quota of a subscription.

**Usage**

	sentinelcli tx subscription quota-add <SUBSCRIPTION_ID> <OWNER_ADDRESS> <BYTES> [FLAGS]

**Example**
	
	// Add a quota of 16 bytes to subscription 1 of address sent1v5ktkl3x9z2ky5uxhqjy487t7x0gl7fjrm0npv.
	sentinelcli tx subscription quota-add 1 sent1v5ktkl3x9z2ky5uxhqjy487t7x0gl7fjrm0npv 16 --from myKeyringWallet --chain-id sentinelhub-2 --node https://rpc.sentinel.co:443 


## Quota Update

**Description**

Update a quota of a subscription.

**Usage**

	sentinelcli tx subscription quota-update <SUBSCRIPTION_ID> <OWNER_ADDRESS> <BYTES> [FLAGS]

**Example**

	// Update the quota of subscription 1 of address sent1v5ktkl3x9z2ky5uxhqjy487t7x0gl7fjrm0npv to 16 bytes.
	sentinelcli tx subscription quota-update 1 sent1v5ktkl3x9z2ky5uxhqjy487t7x0gl7fjrm0npv 16 --from myKeyringWallet --chain-id sentinelhub-2 --node https://rpc.sentinel.co:443


## Subscribe To Node

**Description**

Subscribe to a node.

**Usage**

	sentinelcli tx subscription subscribe-to-node <NODE_ADDRESS> <DEPOSIT_AMOUNT> [FLAGS]	

**Example**
	
	// Subscribe to node sentnode1fj0t8ympfw7k5zt6xsu9ae80k9rva739v2dmp8 using a deposit of 1 DVPN.
	sentinelcli tx subscription subscribe-to-node sentnode1fj0t8ympfw7k5zt6xsu9ae80k9rva739v2dmp8 1000000udvpn --node https://rpc.sentinel.co:443 --from myKeyringWallet --chain-id sentinelhub-2


## Subscribe to Plan

**Description**

Subscribe to a plan. This feature is not implemented yet.

**Usage**

	sentinelcli tx subscription subscribe-to-plan <PLAN> <DENOM> [FLAGS]


## Cancel

**Description**

Cancel a subscription.

**Usage**

	sentinelcli tx subscription cancel <SUBSCRIPTION_ID> [FLAGS]

**Example**
	
	sentinelcli tx subscription cancel 1 --node https://rpc.sentinel.co:443 --from myKeyringWallet --chain-id sentinelhub-2
