import { ValueObject} from '@admin-videos-catalog/core/shared/value-object'
import { Notification } from "@admin-videos-catalog/core/shared/validators";

export abstract class Entity {

  notification: Notification = new Notification();

  abstract get entity_id(): ValueObject;
  abstract toJSON(): any;
}
