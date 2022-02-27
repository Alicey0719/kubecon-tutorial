package v1

import (
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	logf "sigs.k8s.io/controller-runtime/pkg/log"
	"sigs.k8s.io/controller-runtime/pkg/webhook"
)

// log is for logging in this package.
var markdownviewlog = logf.Log.WithName("markdownview-resource")

func (r *MarkdownView) SetupWebhookWithManager(mgr ctrl.Manager) error {
	return ctrl.NewWebhookManagedBy(mgr).
		For(r).
		Complete()
}

// TODO(user): EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!

//+kubebuilder:webhook:path=/mutate-view-k8s-alicey-dev-v1-markdownview,mutating=true,failurePolicy=fail,sideEffects=None,groups=view.k8s.alicey.dev,resources=markdownviews,verbs=create;update,versions=v1,name=mmarkdownview.kb.io,admissionReviewVersions=v1

var _ webhook.Defaulter = &MarkdownView{}

// Default implements webhook.Defaulter so a webhook will be registered for the type
func (r *MarkdownView) Default() {
	markdownviewlog.Info("default", "name", r.Name)

	// TODO(user): fill in your defaulting logic.
}

// TODO(user): change verbs to "verbs=create;update;delete" if you want to enable deletion validation.
//+kubebuilder:webhook:path=/validate-view-k8s-alicey-dev-v1-markdownview,mutating=false,failurePolicy=fail,sideEffects=None,groups=view.k8s.alicey.dev,resources=markdownviews,verbs=create;update,versions=v1,name=vmarkdownview.kb.io,admissionReviewVersions=v1

var _ webhook.Validator = &MarkdownView{}

// ValidateCreate implements webhook.Validator so a webhook will be registered for the type
func (r *MarkdownView) ValidateCreate() error {
	markdownviewlog.Info("validate create", "name", r.Name)

	// TODO(user): fill in your validation logic upon object creation.
	return nil
}

// ValidateUpdate implements webhook.Validator so a webhook will be registered for the type
func (r *MarkdownView) ValidateUpdate(old runtime.Object) error {
	markdownviewlog.Info("validate update", "name", r.Name)

	// TODO(user): fill in your validation logic upon object update.
	return nil
}

// ValidateDelete implements webhook.Validator so a webhook will be registered for the type
func (r *MarkdownView) ValidateDelete() error {
	markdownviewlog.Info("validate delete", "name", r.Name)

	// TODO(user): fill in your validation logic upon object deletion.
	return nil
}
