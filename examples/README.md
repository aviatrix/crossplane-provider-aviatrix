# Aviatrix Provider Examples and Tests

The files in this folder demonstrate how to use the Aviatrix Crossplane Provider.  They describe a Google Cloud Transit Architecture including four VPCs, two peered transit gateways, two attached spokes, 
two associated network domains and a corresponding policy, as well as snat and dnat rules for one of the spokes.  All files can be applied simultaneously to one controller, provided the following prerequisites are met:
 * The Aviatrix Controller is installed and available, with credentials for Crossplane to connect
 * The Aviatrix Provider is installed and configured (see providerconfig folder)
 * A google account named 'aviatrix' on the controller, with all the appropriate permissions (see aviatrix docs for onboarding GCP account)
 * No conflicting resources exist in the gcloud project associated with the aviatrix account
 * A secret like accountsecret.yaml.tmpl exists, with credentials for a valid AWS account (used for account.yaml example only)

Once these prerequisites are met, the examples can be applied by running `kubectl apply -f examples/`.

## Unit Tests

The example files are used as the basis for unit tests located in `internal/provider_test.go`.  These tests use the examples as input, generating terraform files as output and comparing them against the expected output at `internal/test-data`.  For examples which must be repeated (for instance the multiple VPCs which are necessary to support the transit and spoke gateways, which are in turn necessary for segmentation and peering examples), only one instance of the resource type is tested.  The files in the `test-data` can themselves be tested by running `terraform apply` in that folder (after adding the appropriate main.tf.json identifying and configuring the aviatrix provider).  In this manner, we can manually verify that our expected terraform output is correct.  This should be done each time the files in test-data are modified.  

In order to prevent leaking sensitive information, the expected output for resources with sensitive data has been redacted, and those files have been moved to `test-data/partial` to prevent them from being unsuccessfully applied with the other files.  Modifications to these files will need substantial manual testing for verification.