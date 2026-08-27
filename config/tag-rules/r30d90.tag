# Rule file format:
#   - The name of the tag to apply is the name of the file
#   - The tag must exist as label on your gmail account
#   - If you want to use nested tags, use "-" to split levels: "social-twitter" will become "social/twitter"
#   - Every line is a rule.
#   - Any mail will be tagged if any of the rules match (OR)
#   - Every condition on a single rule must match in order to whole rule match (AND)

# This tag is intended to mark mails that are going to
# be automatically "mark as read" 30 days after its reception and
# and  automatically deleted 90 days after its reception

from:no-reply@accounts.google.com
from:noreply@communication.basic-fit.com subject:*cambios*
